<template>
  <div style="margin-bottom: 40px;">
    <VotePopup :proposal="proposal" v-if="store.votePopup"></VotePopup>
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
                <div class="css-79elbk">{{convertSmallToLarge(proposal.total_deposit[0].amount)}} FETF</div>
              </div>
            </div>
          </div>
        </div>
        <div class="button-div">
          <a :href="'https://defund.explorers.guru/proposal/' + proposal.proposal_id" target="_blank" class="sp-button">Details</a>
          <button v-on:click="toggleVotePopup" style="margin-left: 10px;" class="sp-button">Vote</button>
        </div>
    </div>
  </div>
</template>
 
<script>
import { SpTheme } from '@starport/vue';
import VotePopup from './VotePopup.vue';
import { store } from '../store/local/popup.js';
export default {
  name: "GovItem",
  components: {
    SpTheme,
    VotePopup
  },
  props: ["proposal"],
  data(props) {
    return {
      store: store
    }
  },
  methods: {
    toggleVotePopup() {
      store.votePopup = true
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