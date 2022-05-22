<template>
    <div class="cover">
        <div v-on:click="closePopup" class="close-div">
        </div>
        <div class="popup-container">
            <div class="modal-body">
                <div class="details-within">Your Vote (Prop {{proposal.proposal_id}})</div>
                <div style="border-radius: 6px; box-shadow: rgba(50, 50, 93, 0.15) 0px 1px 3px 0px, rgba(0, 0, 0, 0.02) 0px 0px 1px 0px;">
                    <ul class="list-group">
                        <li class="list-group-item" style="cursor: pointer;">
                            <div :proposal="proposal.proposal_id" v-on:click="e => setVoteSelection(e, 'yes')" class="container">
                                <div style="display: grid;grid-template-columns: 0.10fr 0.90fr;" class="justify-content-between align-items-center row">
                                    <div class="col-option" style="padding-right: 0px;">
                                        <div :id="'yes-select-' + proposal.proposal_id" class="voteOption"></div>
                                    </div>
                                    <div class="col">Yes</div>
                                </div>
                            </div>
                        </li>
                        <li class="list-group-item" style="cursor: pointer;">
                            <div :proposal="proposal.proposal_id" v-on:click="e => setVoteSelection(e, 'no')" class="container">
                                <div style="display: grid;grid-template-columns: 0.10fr 0.90fr;" class="justify-content-between align-items-center row">
                                    <div class="col-option" style="padding-right: 0px;">
                                        <div :id="'no-select-' + proposal.proposal_id" class="voteOption"></div>
                                    </div>
                                    <div class="col">No</div>
                                </div>
                            </div>
                        </li>
                        <li class="list-group-item" style="cursor: pointer;">
                            <div :proposal="proposal.proposal_id" v-on:click="e => setVoteSelection(e, 'veto')" class="container">
                                <div style="display: grid;grid-template-columns: 0.10fr 0.90fr;" class="justify-content-between align-items-center row">
                                    <div class="col-option" style="padding-right: 0px;">
                                        <div :id="'veto-select-' + proposal.proposal_id" class="voteOption"></div>
                                    </div>
                                    <div class="col">No With Veto</div>
                                </div>
                            </div>
                        </li>
                        <li class="list-group-item" style="cursor: pointer;">
                            <div :proposal="proposal.proposal_id" v-on:click="e => setVoteSelection(e, 'abstain')" class="container">
                                <div style="display: grid;grid-template-columns: 0.10fr 0.90fr;" class="justify-content-between align-items-center row">
                                    <div class="col-option" style="padding-right: 0px;">
                                        <div :id="'abstain-select-' + proposal.proposal_id" class="voteOption"></div>
                                    </div>
                                    <div class="col">Abstain</div>
                                </div>
                            </div>
                        </li>
                    </ul>
                </div>
                <div class="delegate-button-div">
                    <SpButton v-on:click="closePopup(false)">Cancel</SpButton>
                    <SpButton style="margin-left:10px;">Vote</SpButton>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import { SpTheme, SpButton } from '@starport/vue';
import { store } from '../store/local/popup.js';
import DelegateForm from './DelegateForm.vue';
export default {
    name: "VotePopup",
    components: { SpTheme, SpButton, DelegateForm },
    props: ["proposal"],
    data(props) {
        return {
            store: store
        }
    },
    methods: {
        closePopup: function() {
            if (store.votePopup == false) {
                store.votePopup = true
            } else {
                store.votePopup = false
                store.currentVoteSelection = null
            }
        },
        setVoteSelection(e, value) {
            console.log(e.target.getAttribute("original-title"))
            // Change the current selection back to normal
            if(store.currentVoteSelection) {
                const oldelement = document.getElementById(store.currentVoteSelection + "-select")
                oldelement.style = ""
            }

            // Set the new selected element to highlight
            const newelement = document.getElementById(value + "-select")
            newelement.style = "border-color: #6CE5E8;"

            store.currentVoteSelection = value
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
        box-shadow: 0 15px 35px rgb(50 50 93 / 20%), 0 5px 15px rgb(0 0 0 / 17%);
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
    .row > * {
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