<template>
  <div class="fileview" :class="{selected: isSelected}">
    <span v-for="(p, i) in path"
          @click="setPath(i)"
          class="path">/{{p}}</span>
    <select class="rootSelector" @change="selectRootFn" :value="selectedRoot" title="choose server root">
      <option v-for="root in roots">{{root}}</option>
    </select>
    <div class="reload-files" role="button" @click="reloadFiles" title="update files from server">
      <i :class="{'icon-arrows-cw': !loading, 'icon-spin4': loading, 'animate-spin': loading}"></i></div>
    <file-header></file-header>
    <div class="fileContainer">
      <div v-if="path.length > 1"
           class="file"
           :class="{'focused': focusedFileIndex === -1}"
           @dblclick="changePathToParent">..
      </div>
      <file v-for="(file, index) in files"
            :file="file"
            :class="{'selected': selected(file), 'focused': isFocused(file)}"
            :key="file.name"
            @click.native="selectFile(file, index, $event)"
            @dblclick.native="changePath(file)">
      </file>
    </div>
  </div>
</template>

<script>
  import File from './File.vue'
  import FileHeader from './FileHeader.vue'
  import {Rpc} from '../rpc.js'
  import {EventBus} from '../EventBus'

  export default{
    props: ['roots', 'id'],
    data: () => {
      return {
        files: [],
        loading: false,
        eventListener: null,
        focusedFileIndex: -1
      }
    },
    computed: {
      state () {
        return this.$store.state.states[this.id]
      },
      roots () {
        return this.$store.state.roots
      },
      selectedRoot () {
        return this.$store.state.states[this.id].selectedRoot
      },
      isSelected () {
        return this.$store.state.selectedState === this.id
      },
      path () {
        return [this.selectedRoot].concat(this.$store.state.states[this.id].path)
      },
      pathString () {
        return this.$store.state.states[this.id].path.reduce((acc, p) => {
          return acc + '/' + p
        }, '')
      },
      isViewSelected () {
        return this.$store.state.selectedState === this.id
      },
      focusedFile () {
        if (this.files.length) {
          return this.files[this.focusedFileIndex]
        }
        return ''
      }
    },
    components: {
      File,
      FileHeader
    },
    watch: {
      selectedRoot () {
        this.reloadFiles()
      },
      path () {
        this.reloadFiles()
      }
    },
    methods: {
      selected (file) {
        return this.$store.state.states[this.id] && this.$store.state.states[this.id].selectedFiles.includes(file.name)
      },
      isFocused (file) {
        return this.focusedFile === file
      },
      selectFile (file, index, event) {
        if (event.shiftKey) {
          let begin = Math.min(this.focusedFileIndex, index)
          let end = Math.max(this.focusedFileIndex, index) + 1
          let selectedFiles = this.files
            .slice(begin, end)
            .map(file => file.name)
          this.$store.commit('selectFiles', {stateId: this.id, value: selectedFiles})
          return
        } else if (event.ctrlKey) {
          this.$store.commit('selectFile', {stateId: this.id, value: file.name})
        } else {
          this.$store.commit('selectSingleFile', {stateId: this.id, value: file.name})
        }
        this.focusedFileIndex = index
      },
      selectRootFn (e) {
        this.selectRoot(e.target.value)
      },
      selectRoot (rootName) {
        this.$store.commit('selectRoot', {stateId: this.id, value: rootName})
      },
      changePath (file) {
        if (file.type === 'f') {
          return
        }
        this.$store.commit('changePath', {stateId: this.id, value: file.name})
        this.focusedFileIndex = -1
      },
      changePathToParent () {
        this.$store.commit('changePathToParent', {stateId: this.id})
        this.focusedFileIndex = -1
      },
      setPath (index) {
        this.$store.commit('setPath', {stateId: this.id, value: index})
      },
      reloadFiles () {
        let vm = this
        if (this.loading) {
          return
        }
        this.loading = true
        Rpc.call('ls', [this.selectedRoot + this.pathString])
          .then((response) => {
            let files = response.result.filter((file) => {
              return file.type === 'd'
            }).concat(
              response.result.filter((file) => {
                return file.type === 'f'
              })
            )
            vm.loading = false
            vm.files = files
          })
      }
    },
    created () {
      EventBus.$on('commandFinished', () => {
        this.reloadFiles()
      })
      let vm = this
      this.eventListener = (e) => {
        if (!vm.isViewSelected || vm.$store.state.ui.state !== 'browse') {
          return
        }
        switch (e.key) {
          case 'ArrowUp':
            if (e.ctrlKey || e.shiftKey) {
              vm.$store.commit('selectFile', {stateId: vm.id, value: vm.focusedFile.name})
            }
            let hasParentDirectory = vm.path.length > 1
            let lowerFileIndex = hasParentDirectory
              ? -1
              : 0
            vm.focusedFileIndex = Math.max(vm.focusedFileIndex - 1, lowerFileIndex)
            break
          case 'ArrowDown':
            if (e.ctrlKey || e.shiftKey) {
              vm.$store.commit('selectFile', {stateId: vm.id, value: vm.focusedFile.name})
            }
            vm.focusedFileIndex = Math.min(vm.focusedFileIndex + 1, vm.files.length - 1)
            break
          case 'Tab':
            var newRootSelectionIndex = (vm.roots.indexOf(vm.selectedRoot) + (e.shiftKey ? -1 : 1)) % vm.roots.length
            if (newRootSelectionIndex < 0) {
              newRootSelectionIndex = vm.roots.length - 1
            }
            let newRootSelection = vm.roots[newRootSelectionIndex]
            vm.selectRoot(newRootSelection)
            e.preventDefault()
            break
          case ' ':
            vm.$store.commit('selectFile', {stateId: vm.id, value: vm.focusedFile.name})
            break
          case 'Enter':
            if (vm.focusedFileIndex === -1) {
              vm.changePathToParent()
            } else {
              vm.changePath(vm.focusedFile)
            }
            break
          case 'u':
            vm.reloadFiles()
        }
      }
      window.addEventListener('keydown', this.eventListener)
    },
    destroyed () {
      window.removeEventListener('keydown', this.eventListener)
    }
  }
</script>

<style>
  .fileview {
    width: calc(50% - 30px);
    height: 100%;
    border: 2px solid blue;
    margin-left: 15px;
    margin-top: 10px;
    margin-bottom: 10px;
    padding: 5px;
    background-color: blue;
    float: left;
    user-select: none;
  }

  .fileview.selected {
    border-color: white;
  }

  .reload-files {
    float: right;
    cursor: pointer;
  }

  .reload-files:hover {
    color: white;
  }

  .rootSelector {
    float: right;
  }

  .path:hover {
    cursor: pointer;
    text-decoration: underline;
  }

  .fileHeader {
    cursor: default;
    position: relative;
    width: 100%;
    margin-top: 8px;
    margin-bottom: 5px;
    border-bottom: 1px solid cyan;
  }

  .fileContainer {
    width: 100%;
    overflow-y: auto;
    overflow-x: hidden;
    height: calc(100% - 2em - 12px);
  }

  .file {
    cursor: default;
    position: relative;
    width: 100%;
  }

  .file:hover {
    color: white;
  }

  .file.focused {
    background: gray;
  }

  .file.selected {
    background: yellow;
    color: blue;
  }

  .file.selected.focused {
    background: green;
    color: blue;
  }

  .fileName {
    display: inline-block;
    overflow-x: hidden;
    word-break: break-all;
    width: calc(53% - 2px);
    height: 1em;
    margin: 0;
    white-space: nowrap;
  }

  .fileSize {
    display: inline-block;
    overflow-x: hidden;
    width: calc(12% - 2px);
    height: 1em;
    margin: 0;
    float: right;
    white-space: nowrap;
  }

  .fileModified {
    display: inline-block;
    overflow-x: hidden;
    width: calc(33% - 2px);
    height: 1em;
    margin: 0;
    float: right;
    overflow-y: hidden;
    white-space: nowrap;
  }

  .animate-spin {
    animation-name: spin;
    animation-duration: 2s;
    animation-timing-function: linear;
    animation-delay: initial;
    animation-iteration-count: 3;
    animation-direction: initial;
    animation-fill-mode: initial;
    animation-play-state: initial;
  }
</style>
