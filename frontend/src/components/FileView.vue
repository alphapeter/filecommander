<template>
  <div class="fileview" :class="{selected: isSelected}">
    <template v-if="isInitialized">
      <span v-for="(p, i) in path"
            @click="setPath(i)"
            class="path">/{{p}}</span>
      <select class="rootSelector" @change="selectRoot" :value="selectedRoot">
        <option v-for="root in roots">{{root}}</option>
      </select>
      <file-header></file-header>
      <div class="fileContainer">
        <div v-if="path.length > 1"
             class="file"
             @dblclick="changePathToParent">..</div>
        <file v-for="file in files"
              :file="file"
              :class="{'selected': selected(file)}"
              :key="file.name"
              @click.native="selectFile(file, $event)"
              @dblclick.native="changePath(file)">
        </file>
      </div>
    </template>
  </div>
</template>

<script>
  import File from './File.vue'
  import FileHeader from './FileHeader.vue'
  import { Rpc } from '../rpc.js'
  export default{
    props: ['roots', 'id'],
    data: () => {
      return {
        files: []
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
        return this.isInitialized && this.$store.state.states[this.id].selectedRoot
      },
      isSelected () {
        return this.$store.state.selectedState === this.id
      },
      path () {
        return this.isInitialized && [this.selectedRoot].concat(this.$store.state.states[this.id].path)
      },
      pathString () {
        return this.isInitialized && this.$store.state.states[this.id].path.reduce((acc, p) => {
          return acc + '/' + p
        }, '')
      },
      isInitialized () {
        return this.$store.state.uiState !== 'initializing'
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
      selectFile (file, event) {
        if (event.ctrlKey) {
          this.$store.commit('selectFile', {stateId: this.id, value: file.name})
        } else {
          this.$store.commit('selectSingleFile', {stateId: this.id, value: file.name})
        }
      },
      selectRoot (e) {
        this.$store.commit('selectRoot', {stateId: this.id, value: e.target.value})
      },
      changePath (file) {
        if (file.type === 'f') {
          return
        }
        this.$store.commit('changePath', {stateId: this.id, value: file.name})
      },
      changePathToParent () {
        this.$store.commit('changePathToParent', {stateId: this.id})
      },
      setPath (index) {
        this.$store.commit('setPath', {stateId: this.id, value: index})
      },
      reloadFiles () {
        let vm = this
        Rpc.call('ls', [this.selectedRoot + this.pathString])
          .then((response) => {
            let files = response.result.filter((file) => {
              return file.type === 'd'
            }).concat(
              response.result.filter((file) => {
                return file.type === 'f'
              })
            )
            vm.files = files
          })
      }
    }
  }
</script>

<style>
  .fileview{
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

  .fileview.selected{
    border-color: white;
  }

  .rootSelector{
    float: right;
  }
  .path{

  }
  .path:hover{
    cursor: pointer;
    text-decoration: underline;
  }

  .fileHeader{
    cursor: default;
    position: relative;
    width: 100%;
    margin-top: 8px;
    margin-bottom: 5px;
    border-bottom: 1px solid cyan;
  }
  .fileContainer{
    width: 100%;
    overflow-y: auto;
    overflow-x: hidden;
    height: calc(100% - 2em - 12px);
  }

  .file{
    cursor: default;
    position: relative;
    width: 100%;
  }

  .file:hover{
    background: gray;
  }
  .file.selected{
    background: yellow;
    color: blue;
  }
  .fileName{
    display: inline-block;
    overflow-x: hidden;
    word-break: break-all;
    width: calc(53% - 2px);
    height: 1em;
    margin: 0;
    white-space: nowrap;
  }
  .fileSize{
    display: inline-block;
    overflow-x: hidden;
    width: calc(12% - 2px);
    height: 1em;
    margin: 0;
    float: right;
    white-space: nowrap;
  }
  .fileModified{
    display: inline-block;
    overflow-x: hidden;
    width: calc(33% - 2px);
    height: 1em;
    margin: 0;
    float: right;
    overflow-y: hidden;
    white-space: nowrap;
  }
</style>
