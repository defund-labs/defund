<template>
    <div>
        {{desc}}
    </div>
</template>
 
<script>
import { SpTheme } from '@starport/vue';
export default {
  name: "DescHover",
  components: {
    SpTheme
  },
  props: ["desc"],
  data: () => {
    return {
      desc: this.desc 
    }
  }
};
</script>

<style>
</style>