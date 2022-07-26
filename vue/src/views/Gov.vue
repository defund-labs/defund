<template>
  <div class="child-container">
    <h2 class="title">Proposals</h2>
    <div v-if="getPropLength">
      <div v-for="prop of store.govProps" class="gov-item" v-bind:key="prop.proposal_id">
        <GovItem :proposal="prop">
        </GovItem>
      </div>
      <VotePopup v-if="store.votePopup"></VotePopup>
    </div>
    <div class="no-proposals" v-else>
      No Proposals to View...
    </div>
  </div>
</template>

<script>
import GovItem from '../components/GovItem.vue';
import VotePopup from '../components/VotePopup.vue';
import { computed } from 'vue';
import { useStore } from 'vuex';
import { store } from '../store/local/store.js';
export default {
  name: 'Gov',
  components: { GovItem, VotePopup },
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
  revo-grid[theme=material] revogr-data .rgRow {
      box-shadow: none;
  }
  revo-grid[theme=material] .rowHeaders {
      background-color: white;
  }
  .row {
      margin-left: 0px;
      margin-right: 1px;
      box-sizing: border-box;
      font-size: 0;
  }
</style>
