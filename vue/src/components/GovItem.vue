<template>
  <div style="margin-bottom: 40px;">
    <div class="row">
        <div class="css-1van9nl">
          <div class="three-col-grid">
            <div>
              <h4 class="title-header">#ID</h4>
                <p class="css-79elbk">{{proposal.proposal_id}}</p>
            </div>
            <div>
              <p class="title-header">Title</p>
              <div class="chakra-skeleton css-143twa6">
                <div class="css-79elbk">
                  <span class="css-16vxv8a">
                    <a :href="'https://defund.explorers.guru/proposal/' + proposal.proposal_id" target="_blank">{{proposal.content.title}}</a>
                  </span>
                </div>
              </div>
            </div>
            <div>
              <p class="title-header">Status</p>
              <div class="chakra-skeleton css-143twa6">
                <div class="css-79elbk">
                  <span class="chakra-badge css-1xazz9g">{{proposal.status}}</span>
                </div>
              </div>
            </div>
          </div>
          <div class="three-col-grid">
            <div>
              <p class="title-header">Voting Start</p>
              <div class="chakra-skeleton css-143twa6">
                <div class="css-79elbk">{{extractDate(proposal.voting_start_time)}}</div>
              </div>
            </div>
            <div>
              <p class="title-header">Voting End</p>
              <div class="chakra-skeleton css-143twa6">
                <div class="css-79elbk">{{extractDate(proposal.voting_end_time)}}</div>
              </div>
            </div>
            <div>
              <p class="title-header">Total Deposit</p>
              <div class="chakra-skeleton css-143twa6">
                <div class="css-79elbk"> {{proposal.total_deposit.length > 0 ? convertSmallToLarge(proposal.total_deposit[0].amount) : NaN}} FETF</div>
              </div>
            </div>
          </div>
        </div>
        <div class="button-div">
          <a :href="'https://defund.explorers.guru/proposal/' + proposal.proposal_id" target="_blank" class="sp-button">Details</a>
          <button v-if="proposal.status == 'PROPOSAL_STATUS_VOTING_PERIOD'" :proposal="proposal.proposal_id" v-on:click="e => toggleVotePopup(e)" style="margin-left: 10px;" class="sp-button">Vote</button>
        </div>
    </div>
  </div>
</template>
 
<script>
import { SpTheme } from '@starport/vue';
import { store } from '../store/local/store.js';
export default {
  name: "GovItem",
  components: {
    SpTheme,
  },
  props: ["proposal"],
  data(props) {
    return {
      store: store
    }
  },
  methods: {
    toggleVotePopup(e) {
      store.votePopup = true

      store.currentVoteSelection.proposal_id = e.target.getAttribute("proposal")
    },
    extractDate(string) {
      return string.split("T")[0]
    },
    convertSmallToLarge(string) {
      return String(Number(string)/1000000)
    }
  },
};
</script>

<style>
  .three-col-grid {
    display: grid;
    grid-template-columns: 0.34fr 0.33fr 0.33fr;
    font-size: initial;
  }
  .title-header {
    margin: revert;
    font-size: large;
    font-weight: bold;
  }
  .vote-div {
    text-align: center;
    padding: 20px;
  }
  .button-div {
    margin-top: 20px;
  }
</style>