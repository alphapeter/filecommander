<template>
  <div class="file">
    <div class="fileName">{{file.name}}</div>
    <div class="fileModified">{{modified(file.modified)}}</div>
    <div class="fileSize" v-if="file.type == 'd'">&ltdir&gt</div>
    <div class="fileSize" v-else>{{size(file.size)}}</div>
  </div>
</template>
<script>
  export default{
    props: ['file'],
    methods: {
      size (size) {
        if (size < 1000) {
          return size
        }
        if (size < 1000000) {
          return Math.round(size / 1000) + 'KB'
        }
        if (size < 1000000000) {
          return Math.round(size / 1000000) + 'MB'
        }

        if (size < 1000000000000) {
          return Math.round(size / 1000000000) + 'GB'
        }
        if (size < 1000000000000000) {
          return Math.round(size / 1000000000000) + 'TB'
        }
        return Math.round(size / 1000000000000000) + 'PB'
      },
      modified (date) {
        return new Date(date).toLocaleString()
      }
    }
  }
</script>
<style>

</style>
