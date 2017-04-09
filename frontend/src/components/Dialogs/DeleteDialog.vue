<template>
  <Modal class="overlay">
    <span slot="title">Delete {{fileCount}} file(s)</span>
    <button slot="buttons"
            v-text="'OK'"
            @click="deleteFile">
    </button>
  </Modal>
</template>

<script>
  import Modal from './Modal.vue'
  import { Rpc } from '../../rpc'
  export default {
    components: {
      Modal
    },
    data () {
      return {
        fileCount: 0,
        keypress: null
      }
    },
    methods: {
      deleteFile () {
        this.$store.commit('deleteFileWait')
        let currentState = this.$store.getters.currentState
        let path = this.$store.getters.currentPathString
        let vm = this
        function del (index) {
          Rpc.call('rm', [currentState.selectedRoot + path + '/' + currentState.selectedFiles[index]])
            .then(response => {
              if (response.error) {
                vm.$store.commit('error', response.error)
              } else if (index >= currentState.selectedFiles.length - 1) {
                vm.$store.commit('commandFinished')
              } else {
                del(index + 1)
              }
            })
        }
        del(0)
      }
    },
    created () {
      this.fileCount = this.$store.getters.currentState.selectedFiles.length
      var vm = this
      this.keypress = function (e) {
        if (e.key === 'Enter' && e.target.nodeName !== 'BUTTON') {
          vm.deleteFile()
        }
      }
      window.addEventListener('keyup', this.keypress)
    },
    destroyed () {
      window.removeEventListener('keyup', this.keypress)
    }
  }
</script>

<style>

</style>
