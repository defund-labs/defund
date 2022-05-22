<template>
  <div class="container">
    <h2 class="title">Proposals</h2>
    <div v-if="getPropLength">
      <div v-for="proposal of store.govProps" class="container" v-bind:key="proposal.proposal_id">
        <GovItem :proposal="proposal">
        </GovItem>
      </div>
    </div>
    <div class="no-proposals" v-else>
      No Proposals to View...
    </div>
  </div>
</template>

<script>
import GovItem from '../components/GovItem.vue';
import { computed } from 'vue';
import { useStore } from 'vuex';
import { store } from '../store/local/popup.js';
export default {
  name: 'Gov',
  components: { GovItem },
  data() {
    let $s = useStore()

    let proposals = computed(() => {
      $s.dispatch("cosmos.gov.v1beta1/QueryProposals", {subscribe: false, all: true})
      var allProps = $s["getters"]["cosmos.gov.v1beta1/getProposals"]()["proposals"]
      return allProps
    })

    store.govProps = proposals

    return {
      store,
      proposals
    }
  },
  methods: {
    getPropLength() {
      return this.store.govProps.length > 0
    }
  },
}
</script>

<style>
  .no-proposals {
    font-size: medium;
    text-align: center;
    min-height: 25em;
    display: grid;
    align-items: center;
  }
</style>
