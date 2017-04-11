<template>
  <Modal :disableButtons="true">
    <slot>

    </slot>
  </Modal>
</template>

<script>
  import Modal from './Modal.vue'
  export default {
    components: {
      Modal
    }
  }
</script>

<style>

</style>
