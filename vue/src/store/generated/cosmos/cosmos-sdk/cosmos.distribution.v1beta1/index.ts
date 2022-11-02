import { Client, registry, MissingWalletError } from 'defund-labs-defund-client-ts'

import { Params } from "defund-labs-defund-client-ts/cosmos.distribution.v1beta1/types"
import { ValidatorHistoricalRewards } from "defund-labs-defund-client-ts/cosmos.distribution.v1beta1/types"
import { ValidatorCurrentRewards } from "defund-labs-defund-client-ts/cosmos.distribution.v1beta1/types"
import { ValidatorAccumulatedCommission } from "defund-labs-defund-client-ts/cosmos.distribution.v1beta1/types"
import { ValidatorOutstandingRewards } from "defund-labs-defund-client-ts/cosmos.distribution.v1beta1/types"
import { ValidatorSlashEvent } from "defund-labs-defund-client-ts/cosmos.distribution.v1beta1/types"
import { ValidatorSlashEvents } from "defund-labs-defund-client-ts/cosmos.distribution.v1beta1/types"
import { FeePool } from "defund-labs-defund-client-ts/cosmos.distribution.v1beta1/types"
import { CommunityPoolSpendProposal } from "defund-labs-defund-client-ts/cosmos.distribution.v1beta1/types"
import { DelegatorStartingInfo } from "defund-labs-defund-client-ts/cosmos.distribution.v1beta1/types"
import { DelegationDelegatorReward } from "defund-labs-defund-client-ts/cosmos.distribution.v1beta1/types"
import { CommunityPoolSpendProposalWithDeposit } from "defund-labs-defund-client-ts/cosmos.distribution.v1beta1/types"
import { DelegatorWithdrawInfo } from "defund-labs-defund-client-ts/cosmos.distribution.v1beta1/types"
import { ValidatorOutstandingRewardsRecord } from "defund-labs-defund-client-ts/cosmos.distribution.v1beta1/types"
import { ValidatorAccumulatedCommissionRecord } from "defund-labs-defund-client-ts/cosmos.distribution.v1beta1/types"
import { ValidatorHistoricalRewardsRecord } from "defund-labs-defund-client-ts/cosmos.distribution.v1beta1/types"
import { ValidatorCurrentRewardsRecord } from "defund-labs-defund-client-ts/cosmos.distribution.v1beta1/types"
import { DelegatorStartingInfoRecord } from "defund-labs-defund-client-ts/cosmos.distribution.v1beta1/types"
import { ValidatorSlashEventRecord } from "defund-labs-defund-client-ts/cosmos.distribution.v1beta1/types"


export { Params, ValidatorHistoricalRewards, ValidatorCurrentRewards, ValidatorAccumulatedCommission, ValidatorOutstandingRewards, ValidatorSlashEvent, ValidatorSlashEvents, FeePool, CommunityPoolSpendProposal, DelegatorStartingInfo, DelegationDelegatorReward, CommunityPoolSpendProposalWithDeposit, DelegatorWithdrawInfo, ValidatorOutstandingRewardsRecord, ValidatorAccumulatedCommissionRecord, ValidatorHistoricalRewardsRecord, ValidatorCurrentRewardsRecord, DelegatorStartingInfoRecord, ValidatorSlashEventRecord };

function initClient(vuexGetters) {
	return new Client(vuexGetters['common/env/getEnv'], vuexGetters['common/wallet/signer'])
}

function mergeResults(value, next_values) {
	for (let prop of Object.keys(next_values)) {
		if (Array.isArray(next_values[prop])) {
			value[prop]=[...value[prop], ...next_values[prop]]
		}else{
			value[prop]=next_values[prop]
		}
	}
	return value
}

type Field = {
	name: string;
	type: unknown;
}
function getStructure(template) {
	let structure: {fields: Field[]} = { fields: [] }
	for (const [key, value] of Object.entries(template)) {
		let field = { name: key, type: typeof value }
		structure.fields.push(field)
	}
	return structure
}
const getDefaultState = () => {
	return {
				Params: {},
				ValidatorOutstandingRewards: {},
				ValidatorCommission: {},
				ValidatorSlashes: {},
				DelegationRewards: {},
				DelegationTotalRewards: {},
				DelegatorValidators: {},
				DelegatorWithdrawAddress: {},
				CommunityPool: {},
				
				_Structure: {
						Params: getStructure(Params.fromPartial({})),
						ValidatorHistoricalRewards: getStructure(ValidatorHistoricalRewards.fromPartial({})),
						ValidatorCurrentRewards: getStructure(ValidatorCurrentRewards.fromPartial({})),
						ValidatorAccumulatedCommission: getStructure(ValidatorAccumulatedCommission.fromPartial({})),
						ValidatorOutstandingRewards: getStructure(ValidatorOutstandingRewards.fromPartial({})),
						ValidatorSlashEvent: getStructure(ValidatorSlashEvent.fromPartial({})),
						ValidatorSlashEvents: getStructure(ValidatorSlashEvents.fromPartial({})),
						FeePool: getStructure(FeePool.fromPartial({})),
						CommunityPoolSpendProposal: getStructure(CommunityPoolSpendProposal.fromPartial({})),
						DelegatorStartingInfo: getStructure(DelegatorStartingInfo.fromPartial({})),
						DelegationDelegatorReward: getStructure(DelegationDelegatorReward.fromPartial({})),
						CommunityPoolSpendProposalWithDeposit: getStructure(CommunityPoolSpendProposalWithDeposit.fromPartial({})),
						DelegatorWithdrawInfo: getStructure(DelegatorWithdrawInfo.fromPartial({})),
						ValidatorOutstandingRewardsRecord: getStructure(ValidatorOutstandingRewardsRecord.fromPartial({})),
						ValidatorAccumulatedCommissionRecord: getStructure(ValidatorAccumulatedCommissionRecord.fromPartial({})),
						ValidatorHistoricalRewardsRecord: getStructure(ValidatorHistoricalRewardsRecord.fromPartial({})),
						ValidatorCurrentRewardsRecord: getStructure(ValidatorCurrentRewardsRecord.fromPartial({})),
						DelegatorStartingInfoRecord: getStructure(DelegatorStartingInfoRecord.fromPartial({})),
						ValidatorSlashEventRecord: getStructure(ValidatorSlashEventRecord.fromPartial({})),
						
		},
		_Registry: registry,
		_Subscriptions: new Set(),
	}
}

// initial state
const state = getDefaultState()

export default {
	namespaced: true,
	state,
	mutations: {
		RESET_STATE(state) {
			Object.assign(state, getDefaultState())
		},
		QUERY(state, { query, key, value }) {
			state[query][JSON.stringify(key)] = value
		},
		SUBSCRIBE(state, subscription) {
			state._Subscriptions.add(JSON.stringify(subscription))
		},
		UNSUBSCRIBE(state, subscription) {
			state._Subscriptions.delete(JSON.stringify(subscription))
		}
	},
	getters: {
				getParams: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Params[JSON.stringify(params)] ?? {}
		},
				getValidatorOutstandingRewards: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ValidatorOutstandingRewards[JSON.stringify(params)] ?? {}
		},
				getValidatorCommission: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ValidatorCommission[JSON.stringify(params)] ?? {}
		},
				getValidatorSlashes: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ValidatorSlashes[JSON.stringify(params)] ?? {}
		},
				getDelegationRewards: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.DelegationRewards[JSON.stringify(params)] ?? {}
		},
				getDelegationTotalRewards: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.DelegationTotalRewards[JSON.stringify(params)] ?? {}
		},
				getDelegatorValidators: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.DelegatorValidators[JSON.stringify(params)] ?? {}
		},
				getDelegatorWithdrawAddress: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.DelegatorWithdrawAddress[JSON.stringify(params)] ?? {}
		},
				getCommunityPool: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.CommunityPool[JSON.stringify(params)] ?? {}
		},
				
		getTypeStructure: (state) => (type) => {
			return state._Structure[type].fields
		},
		getRegistry: (state) => {
			return state._Registry
		}
	},
	actions: {
		init({ dispatch, rootGetters }) {
			console.log('Vuex module: cosmos.distribution.v1beta1 initialized!')
			if (rootGetters['common/env/client']) {
				rootGetters['common/env/client'].on('newblock', () => {
					dispatch('StoreUpdate')
				})
			}
		},
		resetState({ commit }) {
			commit('RESET_STATE')
		},
		unsubscribe({ commit }, subscription) {
			commit('UNSUBSCRIBE', subscription)
		},
		async StoreUpdate({ state, dispatch }) {
			state._Subscriptions.forEach(async (subscription) => {
				try {
					const sub=JSON.parse(subscription)
					await dispatch(sub.action, sub.payload)
				}catch(e) {
					throw new Error('Subscriptions: ' + e.message)
				}
			})
		},
		
		
		
		 		
		
		
		async QueryParams({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.CosmosDistributionV1Beta1.query.queryParams()).data
				
					
				commit('QUERY', { query: 'Params', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryParams', payload: { options: { all }, params: {...key},query }})
				return getters['getParams']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryParams API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryValidatorOutstandingRewards({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.CosmosDistributionV1Beta1.query.queryValidatorOutstandingRewards( key.validator_address)).data
				
					
				commit('QUERY', { query: 'ValidatorOutstandingRewards', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryValidatorOutstandingRewards', payload: { options: { all }, params: {...key},query }})
				return getters['getValidatorOutstandingRewards']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryValidatorOutstandingRewards API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryValidatorCommission({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.CosmosDistributionV1Beta1.query.queryValidatorCommission( key.validator_address)).data
				
					
				commit('QUERY', { query: 'ValidatorCommission', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryValidatorCommission', payload: { options: { all }, params: {...key},query }})
				return getters['getValidatorCommission']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryValidatorCommission API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryValidatorSlashes({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.CosmosDistributionV1Beta1.query.queryValidatorSlashes( key.validator_address, query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.CosmosDistributionV1Beta1.query.queryValidatorSlashes( key.validator_address, {...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'ValidatorSlashes', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryValidatorSlashes', payload: { options: { all }, params: {...key},query }})
				return getters['getValidatorSlashes']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryValidatorSlashes API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryDelegationRewards({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.CosmosDistributionV1Beta1.query.queryDelegationRewards( key.delegator_address,  key.validator_address)).data
				
					
				commit('QUERY', { query: 'DelegationRewards', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryDelegationRewards', payload: { options: { all }, params: {...key},query }})
				return getters['getDelegationRewards']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryDelegationRewards API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryDelegationTotalRewards({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.CosmosDistributionV1Beta1.query.queryDelegationTotalRewards( key.delegator_address)).data
				
					
				commit('QUERY', { query: 'DelegationTotalRewards', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryDelegationTotalRewards', payload: { options: { all }, params: {...key},query }})
				return getters['getDelegationTotalRewards']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryDelegationTotalRewards API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryDelegatorValidators({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.CosmosDistributionV1Beta1.query.queryDelegatorValidators( key.delegator_address)).data
				
					
				commit('QUERY', { query: 'DelegatorValidators', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryDelegatorValidators', payload: { options: { all }, params: {...key},query }})
				return getters['getDelegatorValidators']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryDelegatorValidators API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryDelegatorWithdrawAddress({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.CosmosDistributionV1Beta1.query.queryDelegatorWithdrawAddress( key.delegator_address)).data
				
					
				commit('QUERY', { query: 'DelegatorWithdrawAddress', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryDelegatorWithdrawAddress', payload: { options: { all }, params: {...key},query }})
				return getters['getDelegatorWithdrawAddress']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryDelegatorWithdrawAddress API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryCommunityPool({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.CosmosDistributionV1Beta1.query.queryCommunityPool()).data
				
					
				commit('QUERY', { query: 'CommunityPool', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryCommunityPool', payload: { options: { all }, params: {...key},query }})
				return getters['getCommunityPool']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryCommunityPool API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		async sendMsgFundCommunityPool({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.CosmosDistributionV1Beta1.tx.sendMsgFundCommunityPool({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgFundCommunityPool:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgFundCommunityPool:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgWithdrawDelegatorReward({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.CosmosDistributionV1Beta1.tx.sendMsgWithdrawDelegatorReward({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgWithdrawDelegatorReward:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgWithdrawDelegatorReward:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgSetWithdrawAddress({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.CosmosDistributionV1Beta1.tx.sendMsgSetWithdrawAddress({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSetWithdrawAddress:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgSetWithdrawAddress:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgWithdrawValidatorCommission({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.CosmosDistributionV1Beta1.tx.sendMsgWithdrawValidatorCommission({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgWithdrawValidatorCommission:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgWithdrawValidatorCommission:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		
		async MsgFundCommunityPool({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.CosmosDistributionV1Beta1.tx.msgFundCommunityPool({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgFundCommunityPool:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgFundCommunityPool:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgWithdrawDelegatorReward({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.CosmosDistributionV1Beta1.tx.msgWithdrawDelegatorReward({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgWithdrawDelegatorReward:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgWithdrawDelegatorReward:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgSetWithdrawAddress({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.CosmosDistributionV1Beta1.tx.msgSetWithdrawAddress({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSetWithdrawAddress:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgSetWithdrawAddress:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgWithdrawValidatorCommission({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.CosmosDistributionV1Beta1.tx.msgWithdrawValidatorCommission({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgWithdrawValidatorCommission:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgWithdrawValidatorCommission:Create Could not create message: ' + e.message)
				}
			}
		},
		
	}
}
