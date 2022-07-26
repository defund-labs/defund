<template>
  <div class="container">
    <header class="funds-header-create">
      <h2 class="title">Create Fund</h2>
    </header>
    <CreateFund></CreateFund>
  </div>
</template>
 
<script>
import { VGrid, VGridVueTemplate } from "@revolist/vue3-datagrid";
import { computed } from 'vue';
import { useStore } from 'vuex';
import { SpTheme, SpButton } from '@starport/vue';
import CreateFund from '../components/CreateFund.vue'

export default {
  name: "CreateFunds",
  data() {
    let $s = useStore()

    return {
    };
  },
  components: {
    VGrid, SpTheme, SpButton, CreateFund
  },
};
</script>

<style>
  .funds-header-create {
    display: flex;
    align-content: center;
    align-items: center;
    justify-content: center;
  }
</style>
