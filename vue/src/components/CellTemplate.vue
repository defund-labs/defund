<template>
    <span>
        {{desc}}
    </span>
</template>
 
<script>
import { SpTheme } from '@starport/vue';
export default {
  name: "CellTemplate",
  components: {
    SpTheme
  },
  data: () => {
    return {
      desc: this.desc 
    }
  }
};
</script>

<style>
</style>