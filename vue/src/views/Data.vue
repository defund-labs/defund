<template>
  <div class="container">
    <!-- <SpCrud store-name="defundhub.defund.etf" item-name="Fund" /> -->
    <CreateFund
      store-name="defundhub.defund.etf"
      item-name="Fund"
      command-name="/sendMsgCreateFund"
    />
  </div>
</template>

<script>
  import CreateFund from '../components/CreateFund.vue'
  import { computed } from 'vue'
  import { useStore } from 'vuex'

  export default {
    name: 'Data',

    components: { CreateFund },

    setup() {
      // store
      let $s = useStore()

      // computed
      let address = computed(() => $s.getters['common/wallet/address'])

      return {
        address
      }
    }
  }
</script>