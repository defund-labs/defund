<template>
    <form class="">
        <div class="form-group">
            <div class="input-group">
                <div class="amt-input">
                    <input v-on:input="onInputDelegateChange" v-on:change="onInputDelegateChange" id="amt-input-undelegate" name="amount" type="number" class="input input-cust" aria-invalid="false">
                    <SpButton v-on:click="setAllValue" type="button" class="button max-button">ALL</SpButton>
                </div>
            </div>
        </div>
        <div class="button-actions-div">
            <SpButton v-on:click="toggleInputUndelegate" type="button">Back</SpButton>
            <SpButton v-on:click="submitUndelegate" :disabled="store.valueUndelegate" style="margin-left:10px;">Undelegate</SpButton>
        </div>
    </form>
</template>

<script>
import { store } from '../store/local/popup.js';
import { SpButton } from '@starport/vue'
import { computed, onMounted } from 'vue'
import { useStore } from 'vuex'
import _ from 'lodash';
export default {
    name: "UndelegateForm",
    components: { SpButton },
    props: ["validator"],
    setup(props, { emit }) {
        let $s = useStore()

        let address = computed(() => {
            return $s.getters['common/wallet/address']
        })

        var balances = computed(() => {
            onMounted(async () => {
                if (address.value) {
                    $s.dispatch('cosmos.bank.v1beta1/QueryAllBalances', {
                        params: { address: address.value },
                        options: { subscribe: true }
                    })
                }
            })

            return $s.getters['cosmos.bank.v1beta1/getAllBalances']({
                params: { address: address.value },
            })
        })

        var balance = null

        if (address.value) {

            balances = JSON.parse(JSON.stringify(balances.value)).balances

            balance = _.filter(balances, function(o) { return o.denom == "ufetf" })[0]

        }

        //Create send delegate msg function
        let submitUndelegate = async () => {
            let address = computed(() => {
                return $s.getters['common/wallet/address']
            })
            const amtInput = document.getElementById('amt-input-undelegate')
            const amount = amtInput.value * 1000000
            $s.dispatch("cosmos.staking.v1beta1/sendMsgDelegate", {
                value: { delegator_address: address.value,
                validator_address: store.currentValidator.operator_address,
                amount: {
                    denom: "ufetf",
                    amount: String(amount)
                }}
            })
            emit('close')
        }

        return {
            balance: balance,
            store: store,
            submitUndelegate
        }
    },
    methods: {
        toggleInput: function() {
            if (store.delegateInput == false) {
                store.delegateInput = true
            } else {
                store.delegateInput = false
                store.delegateInput = true
            }
        },
        toggleInputUndelegate: function() {
            if (store.undelegateInput == false) {
                store.undelegateInput = true
                store.showDelegateButtons = false
            } else {
                store.undelegateInput = false
                store.undelegateInput = true
                store.showDelegateButtons = true
            }
        },
        setAllValue: function() {
            const amtInput = document.getElementById('amt-input-undelegate')
            amtInput.value = this.balance.amount/1000000
            if (amtInput.value != "" && Number(amtInput.value) > 0) {
                store.valueUndelegate = false
            } else {
                store.valueUndelegate = true
            }
        },
        onInputDelegateChange: function(element) {
            if (element.target.value != "" && Number(element.target.value) > 0) {
                store.valueUndelegate = false
            } else {
                store.valueUndelegate = true
            }
        }
    },
}
</script>

<style>
    .max-button {
        right: 80px;
        height: 40px;
        max-width: 70px;
    }
    .amt-input {
        display: -webkit-inline-box;
        -webkit-box-align: center;
        width: inherit;
    }
    input::-webkit-outer-spin-button,
    input::-webkit-inner-spin-button {
    -webkit-appearance: none;
    margin: 0;
    }
    input[type=number] {
    -moz-appearance: textfield;
    }
    .input-group {
        position: relative;
        display: flex;
        flex-wrap: nowrap;
        width: 100%;
        flex-direction: row;
        align-content: center;
        justify-content: center;
        align-items: stretch;
    }
    .input-group-text {
        display: flex;
        align-items: center;
        padding: 0.625rem 0.75rem;
        margin-bottom: 0;
        font-weight: bold;
        line-height: 1.5;
        text-align: center;
        white-space: nowrap;
        background-color: #fff;
        border: 1px solid #dee2e6;
        border-radius: 0.25rem;
        width: 80px;
        justify-content: center;
    }
    .button-actions-div{
        margin-top: 15px;
        float: right;
    }
    .input-cust {
        width: 100%;
        display: inherit;
    }
</style>