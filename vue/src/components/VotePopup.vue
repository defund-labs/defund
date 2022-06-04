<template>
    <div class="cover">
        <div v-on:click="closePopup" class="close-div">
        </div>
        <div v-if="proposal" class="popup-container">
            <div class="modal-body">
                <div class="details-within">Your Vote (Prop {{proposal.proposal_id}})</div>
                <div style="border-radius: 6px;">
                    <ul class="list-group">
                        <li v-on:click="e => setVoteSelection(e, 'yes')" class="list-group-item" style="cursor: pointer;">
                            <div class="container">
                                <div style="display: grid;grid-template-columns: 0.10fr 0.90fr;" class="justify-content-between align-items-center row-option">
                                    <div class="col-option" style="padding-right: 0px;">
                                        <div id="yes-option" class="voteOption"></div>
                                    </div>
                                    <div class="col">Yes</div>
                                </div>
                            </div>
                        </li>
                        <li v-on:click="e => setVoteSelection(e, 'no')" class="list-group-item" style="cursor: pointer;">
                            <div class="container">
                                <div style="display: grid;grid-template-columns: 0.10fr 0.90fr;" class="justify-content-between align-items-center row-option">
                                    <div class="col-option" style="padding-right: 0px;">
                                        <div id="no-option" class="voteOption"></div>
                                    </div>
                                    <div class="col">No</div>
                                </div>
                            </div>
                        </li>
                        <li v-on:click="e => setVoteSelection(e, 'veto')" class="list-group-item" style="cursor: pointer;">
                            <div class="container">
                                <div style="display: grid;grid-template-columns: 0.10fr 0.90fr;" class="justify-content-between align-items-center row-option">
                                    <div class="col-option" style="padding-right: 0px;">
                                        <div id="veto-option" class="voteOption"></div>
                                    </div>
                                    <div class="col">No With Veto</div>
                                </div>
                            </div>
                        </li>
                        <li v-on:click="e => setVoteSelection(e, 'abstain')" class="list-group-item" style="cursor: pointer;">
                            <div class="container">
                                <div style="display: grid;grid-template-columns: 0.10fr 0.90fr;" class="justify-content-between align-items-center row-option">
                                    <div class="col-option" style="padding-right: 0px;">
                                        <div id="abstain-option" class="voteOption"></div>
                                    </div>
                                    <div class="col">Abstain</div>
                                </div>
                            </div>
                        </li>
                    </ul>
                </div>
                <div class="delegate-button-div">
                    <SpButton v-on:click="closePopup(false)">Cancel</SpButton>
                    <SpButton v-on:click="submitVote" style="margin-left:10px;">Vote</SpButton>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import { computed } from 'vue';
import { SpTheme, SpButton } from '@starport/vue';
import { store } from '../store/local/store.js';
import { useStore } from 'vuex';
export default {
    name: "VotePopup",
    components: { SpTheme, SpButton },
    data() {
        const $s = useStore();

        let proposal = computed(() => {
            $s.dispatch("cosmos.gov.v1beta1/QueryProposal", { params:{ proposal_id: store.currentVoteSelection.proposal_id } , subscribe: false, all: false})
            var prop = $s["getters"]["cosmos.gov.v1beta1/getProposal"]({ params: { proposal_id: store.currentVoteSelection.proposal_id }})["proposal"]
            return prop
        })

        let creator = computed(() => $s.getters['common/wallet/address'])

        const submitVote = async () => {
            const voteMap = {
                "yes": 1,
                "no": 3,
                "abstain": 2,
                "veto": 4
            }
            const value = {
                proposal_id: this.store.currentVoteSelection.proposal_id,
                voter: this.creator,
                option: voteMap[this.store.currentVoteSelection.current_vote]
            }

            this.store.sendingTx = true
            this.store.showTxStatus = true

            const res = await $s.dispatch("cosmos.gov.v1beta1/sendMsgVote", {
                value: value,
                fee: [{
                    amount: "200000",
                    denom: "ufetf"
                }],
                memo: ""
            })

            if(res.code == 0) {
                this.store.sendingTx= false
                this.store.showTxSuccess = true 
                this.store.showTxFail = false
                this.store.lastTxHash = res.transactionHash
            } else {
                this.store.sendingTx= false
                this.store.showTxSuccess = false
                this.store.showTxFail = true
                this.store.lastTxHash = res.transactionHash
                this.store.lastTxLog = res.rawLog
            }
            return res
        }

        return {
            store: store,
            proposal,
            creator,
            submitVote
        }
    },
    methods: {
        closePopup() {
            if (store.votePopup == false) {
                store.votePopup = true
            } else {
                store.votePopup = false
                const oldSelect = document.getElementById(store.currentVoteSelection.current_vote + "-option")
                oldSelect.style = ""
                store.currentVoteSelection = {
                    proposal_id: null,
                    current_vote: null
                }
            }
        },
        setVoteSelection(e, selection) {
            if (store.currentVoteSelection.current_vote) {
                const oldSelect = document.getElementById(store.currentVoteSelection.current_vote + "-option")
                oldSelect.style = ""
            }

            const newSelect = document.getElementById(selection + "-option")
            newSelect.style = "border-color: #6CE5E8;"
            store.currentVoteSelection.current_vote = selection
        }
    },
}
</script>

<style>
    .cover {
        position: fixed;
        top: 0;
        right: 0;
        bottom: 0;
        left: 0;
        z-index: 100;
        overflow: hidden;
        outline: 0;
        background-color: rgba(0, 0, 0, 0.185);
        display: flex;
        -webkit-box-pack: center;
        -webkit-justify-content: center;
        -ms-flex-pack: center;
        justify-content: center;
        -webkit-box-align: center;
        -webkit-align-items: center;
        -ms-flex-align: center;
        align-items: center;
    }
    .popup-container {
        font-size: medium;
        z-index: 1000;
        padding: 1.5rem;
        min-width: 500px;
        background-color: white;
        border: 0 solid rgba(0,0,0,.2);
        border-radius: 0.4375rem;
    }
    .close-div{
        position: absolute;
        width: 100vw;
        height: 100vh;
    }
    .modal-body {
        position: relative;
        flex: 1 1 auto;
        padding: 1.5rem;
    }
    .align-items-center {
        align-items: center !important;
    }
    .media {
        display: flex;
        align-items: flex-start;
    }
    .media-2 {
        display: flex;
        align-items: flex-start;
        margin-left: 0px;
    }
    .pic-size {
        width: 74px;
        height: 74px;
    }
    .pic {
        color: #fff;
        display: inline-flex;
        align-items: center;
        justify-content: center;
        font-size: 1rem;
        border-radius: 0.375rem;
        height: 75px;
        width: 75px;
    }
    .img {
        border-radius: 100px;
        height: -webkit-fill-available;
    }
    .details-div {
        margin-top: 15px;
        margin-bottom: 15px;
    }
    .desc-div {
        margin-top: 15px;
    }
    .details-within {
        font-weight: bold;
        margin-bottom: 10px;
    }
    .delegate-button-div {
        float: right;
        margin-top: 15px;
    }
    .list-group-item {
        position: relative;
        display: block;
        padding-top: 20px;
        padding-bottom: 20px;
        margin-bottom: -1px;
        background-color: #fff;
        border: 1px solid #e9ecef;
    }
    .list-group {
        display: flex;
        flex-direction: column;
        padding-left: 0;
        margin-bottom: 0;
    }
    .row {
        font-size: medium;
        border-color: rgba(0, 0, 0, 0.027);
        border-width: 2px;
        border-bottom-style: solid;
        padding-bottom: 25px;
    }
    .row-option {
        font-size: medium;
    }
    .voteOption {
        width: 15px;
        height: 15px;
        font-size: 15px;
        border-radius: 50%;
        border-width: 5px;
        border-style: solid;
        cursor: pointer;
        border-color: #b3b3b3;
    }
    .col-option {
        display: grid;
        grid-template-columns: 0.10fr 0.90fr;
    }
</style>