<template>
  <div id="app">
    <div class="fileviews">
      <file-view id=0 @click.native="selectView('0', $event)"></file-view>
      <file-view id=1 @click.native="selectView('1', $event)"></file-view>
    </div >
    <commands></commands>
    <wait-dialog v-if="$store.state.uiState === 'initializing'">loading...</wait-dialog>
    <wait-dialog v-if="$store.state.uiState === 'rename-wait'">renaming file</wait-dialog>
    <wait-dialog v-if="$store.state.uiState === 'mkdir-wait'">creating new directory</wait-dialog>

    <progress-dialog v-if="$store.state.uiState === 'copy-wait'">copying files...</progress-dialog>
    <progress-dialog v-if="$store.state.uiState === 'move-wait'">moving files...</progress-dialog>
    <progress-dialog v-if="$store.state.uiState === 'delete-file-wait'">deleting files...</progress-dialog>

    <rename-dialog v-if="$store.state.uiState === 'rename'">renaming file</rename-dialog>
    <mkdir-dialog v-if="$store.state.uiState === 'mkdir'"></mkdir-dialog>
    <error-dialog v-if="$store.state.uiState === 'error'"></error-dialog>
    <delete-dialog v-if="$store.state.uiState === 'delete-file'"></delete-dialog>
  </div>
</template>

<script>
  import Commands from './components/Commands.vue'
  import FileView from './components/FileView.vue'
  import Header from './components/Header.vue'
  import WaitDialog from './components/Dialogs/WaitDialog.vue'
  import RenameDialog from './components/Dialogs/RenameDialog.vue'
  import MkdirDialog from './components/Dialogs/MkdirDialog.vue'
  import ErrorDialog from './components/Dialogs/ErrorDialog.vue'
  import DeleteDialog from './components/Dialogs/DeleteDialog.vue'
  import ProgressDialog from './components/Dialogs/ProgressDialog.vue'

  export default {
    name: 'app',
    components: {
      Commands,
      FileView,
      'appHeader': Header,
      WaitDialog,
      RenameDialog,
      MkdirDialog,
      DeleteDialog,
      ErrorDialog,
      ProgressDialog
    },
    data () {
      return {
        roots: [],
        selectedView: null
      }
    },
    methods: {
      selectView (viewId, event) {
        if (event.target.nodeName === 'SELECT') {
          return
        }
        this.$store.commit('selectView', viewId)
      }
    },
    created () {
      this.$store.dispatch('init')
    }
  }
</script>

<style>
  #app {
    height:100%;
  }
  .fileviews{
    position: relative;
    width: 100%;
    height: calc(100% - 61px);
  }
</style>
