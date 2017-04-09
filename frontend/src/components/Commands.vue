<template>
  <div class="commands">
    <button id="rename"
            class="icon-pencil-squared"
            :disabled="buttonsDisabled || multipleFilesSelected"
            @click="rename"
            v-text="'Rename (R)'">

    </button>
    <button id="mkdir"
            class="icon-folder-empty-1"
            @click="mkdir"
            v-text="'New directory (N)'">
    </button>
    <button id="copy"
            class="icon-clone"
            :disabled="buttonsDisabled"
            @click="copy"
            v-text="'Copy (C)'">
    </button>
    <button id="move"
            class="icon-exchange"
            :disabled="buttonsDisabled"
            @click="move"
            v-text="'Move (O)'">
    </button>
    <button id="delete"
            class="icon-trash-empty"
            :disabled="buttonsDisabled"
            @click="deleteFile"
            v-text="'Delete (D)'">
    </button>
  </div>
</template>

<script>
  import { Rpc } from '../rpc.js'
  export default {
    computed: {
      buttonsDisabled () {
        return this.$store.state.states[this.$store.state.selectedState] && this.$store.state.states[this.$store.state.selectedState].selectedFiles.length === 0
      },
      multipleFilesSelected () {
        return this.$store.state.states[this.$store.state.selectedState] && this.$store.state.states[this.$store.state.selectedState].selectedFiles.length > 1
      }
    },
    data: () => {
      return {
        eventListener: null
      }
    },
    methods: {
      mkdir () {
        this.$store.commit('mkdir')
      },
      copy () {
        if (this.buttonsDisabled) {
          return
        }
        this.$store.commit('copyWait')
        this.executeBinaryCommand('cp')
      },
      move () {
        if (this.buttonsDisabled) {
          return
        }
        this.$store.commit('moveWait')
        this.executeBinaryCommand('mv')
      },
      deleteFile () {
        if (this.buttonsDisabled) {
          return
        }
        this.$store.commit('deleteFile')
      },
      rename () {
        if (this.buttonsDisabled || this.multipleFilesSelected) {
          return
        }
        this.$store.commit('rename')
      },
      executeBinaryCommand (command) {
        let currentState = this.$store.getters.currentState
        let currentPath = this.$store.getters.currentPathString
        let otherState = this.$store.getters.otherState
        let otherPath = this.$store.getters.otherPathString

        let vm = this
        function run (index) {
          Rpc.call(command,
            [currentState.selectedRoot + currentPath + '/' + currentState.selectedFiles[index],
              otherState.selectedRoot + otherPath + '/' + currentState.selectedFiles[index]])
            .then(response => {
              if (response.error) {
                vm.$store.commit('error', response.error)
              } else if (index >= currentState.selectedFiles.length - 1) {
                vm.$store.commit('commandFinished')
              } else {
                run(index + 1)
              }
            })
        }
        run(0)
      }
    },
    created () {
      let vm = this
      this.eventListener = (e) => {
        if (vm.$store.state.uiState !== 'browse') {
          return
        }
        switch (e.key) {
          case 'r':
            this.rename()
            break
          case 'Insert':
          case 'n':
            this.mkdir()
            break
          case 'c':
            this.copy()
            break
          case 'm':
            this.move()
            break
          case 'Delete':
          case 'd':
            this.deleteFile()
            break
          default:
            break
        }
        if (e.key === 'Escape') {
          vm.$store.state.uiState = 'browse'
        }
      }
      window.addEventListener('keyup', this.eventListener)
    },
    destroyed () {
      if (this.disableButtons) {
        return
      }
      window.removeEventListener('keyup', this.eventListener)
    }
  }
</script>

<style>
  .commands{
    clear: both;
  }
</style>
