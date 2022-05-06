<template>
    <div class="cover">
        <div v-on:click="closePopup" class="close-div">
        </div>
        <div class="popup-container">
            <div class="modal-1AN7AqKNS_NB76-16UlB44 modal-body">
                <div>
                    <div class="align-items-center media-2">
                        <div class="media" style="flex-direction: column;">
                            <span class="name text-lg">
                                <strong>{{store.currentValidator["description.moniker"]}}</strong>
                            </span>
                            <span class="text-sm">Commission - {{store.currentValidator["commission.commission_rates.rate"]}}</span>
                        </div>
                    </div>
                    <div class="details-div">
                        <div v-if="store.currentValidator['description.website']" class="website-div">
                            <div class="details-within">Website</div>
                            <a :href="'https://www.' + store.currentValidator['description.website']" target="_blank" rel="noopener noreferrer">{{store.currentValidator["description.website"]}}</a>
                        </div>
                        <div v-if="store.currentValidator['description.details']" class="desc-div">
                            <div class="details-within">Description</div>
                            <div>{{store.currentValidator["description.details"]}}</div>
                        </div>
                        <div class="desc-div">
                            <div class="details-within">My Delegations</div>
                            <p>{{delegations ? String(Number(delegations.balance.amount)/1000000) : 0}} FETF</p>
                        </div>
                    </div>
                    <div v-if="!store.delegateInput" class="delegate-button-div">
                        <SpButton v-on:click="toggleInput(false)">Delegate</SpButton>
                        <SpButton v-on:click="toggleInput(true)" style="margin-left:10px;">Undelegate</SpButton>
                    </div>
                    <div v-if="store.delegateInput">
                        <DelegateForm :delegations="delegations"></DelegateForm>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import { SpTheme, SpButton } from '@starport/vue';
import { useStore } from 'vuex';
import { computed } from 'vue';
import { store } from '../store/local/popup.js';
import DelegateForm from './DelegateForm.vue'
export default {
    name: "StakePopup",
    components: { SpTheme, SpButton, DelegateForm },
    setup(props) {
        let $s = useStore()

        let address = computed(() => {
            return $s.getters['common/wallet/address']
        })

        let delegations = computed(() => {
            $s.dispatch("cosmos.staking.v1beta1/QueryDelegation", { params: {
                validator_addr: store.currentValidator.operator_address,  delegator_addr: address.value
            }, subscribe: true, all: false })
            var allDelegations = $s.getters["cosmos.staking.v1beta1/getDelegation"]({
                params: { validator_addr: store.currentValidator.operator_address,  delegator_addr: address.value },
            })
            return allDelegations
        })

        return {
            delegations: JSON.parse(JSON.stringify(delegations.value)).delegation_response,
            store: store
        }
    },
    methods: {
        closePopup: function() {
            if (store.stakePopup == false) {
                store.stakePopup = true
            } else {
                store.stakePopup = false
                store.valueDelegate = true
                store.valueUndelegate = true
                store.currentValidator = null
                store.showDelegateButtons = false
            }
        },
        toggleInput: function(undelegate = false) {
            // If undelegate button flag is set, set input status store as undelegate
            if (undelegate) {
                store.undelegate = true
            } else {
                store.undelegate = false
            }
            if (store.delegateInput == true) {
                store.delegateInput = false
            }
            if (store.delegateInput == false) {
                store.delegateInput = true
            } else {
                store.delegateInput = false
            }
        },
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
</style>
