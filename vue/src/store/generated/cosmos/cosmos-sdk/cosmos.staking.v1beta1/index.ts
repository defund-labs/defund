import { Client, registry, MissingWalletError } from 'defund-labs-defund-client-ts'

import { StakeAuthorization } from "defund-labs-defund-client-ts/cosmos.staking.v1beta1/types"
import { StakeAuthorization_Validators } from "defund-labs-defund-client-ts/cosmos.staking.v1beta1/types"
import { LastValidatorPower } from "defund-labs-defund-client-ts/cosmos.staking.v1beta1/types"
import { HistoricalInfo } from "defund-labs-defund-client-ts/cosmos.staking.v1beta1/types"
import { CommissionRates } from "defund-labs-defund-client-ts/cosmos.staking.v1beta1/types"
import { Commission } from "defund-labs-defund-client-ts/cosmos.staking.v1beta1/types"
import { Description } from "defund-labs-defund-client-ts/cosmos.staking.v1beta1/types"
import { Validator } from "defund-labs-defund-client-ts/cosmos.staking.v1beta1/types"
import { ValAddresses } from "defund-labs-defund-client-ts/cosmos.staking.v1beta1/types"
import { DVPair } from "defund-labs-defund-client-ts/cosmos.staking.v1beta1/types"
import { DVPairs } from "defund-labs-defund-client-ts/cosmos.staking.v1beta1/types"
import { DVVTriplet } from "defund-labs-defund-client-ts/cosmos.staking.v1beta1/types"
import { DVVTriplets } from "defund-labs-defund-client-ts/cosmos.staking.v1beta1/types"
import { Delegation } from "defund-labs-defund-client-ts/cosmos.staking.v1beta1/types"
import { UnbondingDelegation } from "defund-labs-defund-client-ts/cosmos.staking.v1beta1/types"
import { UnbondingDelegationEntry } from "defund-labs-defund-client-ts/cosmos.staking.v1beta1/types"
import { RedelegationEntry } from "defund-labs-defund-client-ts/cosmos.staking.v1beta1/types"
import { Redelegation } from "defund-labs-defund-client-ts/cosmos.staking.v1beta1/types"
import { Params } from "defund-labs-defund-client-ts/cosmos.staking.v1beta1/types"
import { DelegationResponse } from "defund-labs-defund-client-ts/cosmos.staking.v1beta1/types"
import { RedelegationEntryResponse } from "defund-labs-defund-client-ts/cosmos.staking.v1beta1/types"
import { RedelegationResponse } from "defund-labs-defund-client-ts/cosmos.staking.v1beta1/types"
import { Pool } from "defund-labs-defund-client-ts/cosmos.staking.v1beta1/types"


export { StakeAuthorization, StakeAuthorization_Validators, LastValidatorPower, HistoricalInfo, CommissionRates, Commission, Description, Validator, ValAddresses, DVPair, DVPairs, DVVTriplet, DVVTriplets, Delegation, UnbondingDelegation, UnbondingDelegationEntry, RedelegationEntry, Redelegation, Params, DelegationResponse, RedelegationEntryResponse, RedelegationResponse, Pool };

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
				Validators: {},
				Validator: {},
				ValidatorDelegations: {},
				ValidatorUnbondingDelegations: {},
				Delegation: {},
				UnbondingDelegation: {},
				DelegatorDelegations: {},
				DelegatorUnbondingDelegations: {},
				Redelegations: {},
				DelegatorValidators: {},
				DelegatorValidator: {},
				HistoricalInfo: {},
				Pool: {},
				Params: {},
				
				_Structure: {
						StakeAuthorization: getStructure(StakeAuthorization.fromPartial({})),
						StakeAuthorization_Validators: getStructure(StakeAuthorization_Validators.fromPartial({})),
						LastValidatorPower: getStructure(LastValidatorPower.fromPartial({})),
						HistoricalInfo: getStructure(HistoricalInfo.fromPartial({})),
						CommissionRates: getStructure(CommissionRates.fromPartial({})),
						Commission: getStructure(Commission.fromPartial({})),
						Description: getStructure(Description.fromPartial({})),
						Validator: getStructure(Validator.fromPartial({})),
						ValAddresses: getStructure(ValAddresses.fromPartial({})),
						DVPair: getStructure(DVPair.fromPartial({})),
						DVPairs: getStructure(DVPairs.fromPartial({})),
						DVVTriplet: getStructure(DVVTriplet.fromPartial({})),
						DVVTriplets: getStructure(DVVTriplets.fromPartial({})),
						Delegation: getStructure(Delegation.fromPartial({})),
						UnbondingDelegation: getStructure(UnbondingDelegation.fromPartial({})),
						UnbondingDelegationEntry: getStructure(UnbondingDelegationEntry.fromPartial({})),
						RedelegationEntry: getStructure(RedelegationEntry.fromPartial({})),
						Redelegation: getStructure(Redelegation.fromPartial({})),
						Params: getStructure(Params.fromPartial({})),
						DelegationResponse: getStructure(DelegationResponse.fromPartial({})),
						RedelegationEntryResponse: getStructure(RedelegationEntryResponse.fromPartial({})),
						RedelegationResponse: getStructure(RedelegationResponse.fromPartial({})),
						Pool: getStructure(Pool.fromPartial({})),
						
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
				getValidators: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Validators[JSON.stringify(params)] ?? {}
		},
				getValidator: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Validator[JSON.stringify(params)] ?? {}
		},
				getValidatorDelegations: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ValidatorDelegations[JSON.stringify(params)] ?? {}
		},
				getValidatorUnbondingDelegations: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ValidatorUnbondingDelegations[JSON.stringify(params)] ?? {}
		},
				getDelegation: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Delegation[JSON.stringify(params)] ?? {}
		},
				getUnbondingDelegation: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.UnbondingDelegation[JSON.stringify(params)] ?? {}
		},
				getDelegatorDelegations: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.DelegatorDelegations[JSON.stringify(params)] ?? {}
		},
				getDelegatorUnbondingDelegations: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.DelegatorUnbondingDelegations[JSON.stringify(params)] ?? {}
		},
				getRedelegations: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Redelegations[JSON.stringify(params)] ?? {}
		},
				getDelegatorValidators: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.DelegatorValidators[JSON.stringify(params)] ?? {}
		},
				getDelegatorValidator: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.DelegatorValidator[JSON.stringify(params)] ?? {}
		},
				getHistoricalInfo: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.HistoricalInfo[JSON.stringify(params)] ?? {}
		},
				getPool: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Pool[JSON.stringify(params)] ?? {}
		},
				getParams: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Params[JSON.stringify(params)] ?? {}
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
			console.log('Vuex module: cosmos.staking.v1beta1 initialized!')
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
		
		
		
		 		
		
		
		async QueryValidators({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.CosmosStakingV1Beta1.query.queryValidators(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.CosmosStakingV1Beta1.query.queryValidators({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'Validators', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryValidators', payload: { options: { all }, params: {...key},query }})
				return getters['getValidators']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryValidators API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryValidator({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.CosmosStakingV1Beta1.query.queryValidator( key.validator_addr)).data
				
					
				commit('QUERY', { query: 'Validator', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryValidator', payload: { options: { all }, params: {...key},query }})
				return getters['getValidator']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryValidator API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryValidatorDelegations({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.CosmosStakingV1Beta1.query.queryValidatorDelegations( key.validator_addr, query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.CosmosStakingV1Beta1.query.queryValidatorDelegations( key.validator_addr, {...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'ValidatorDelegations', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryValidatorDelegations', payload: { options: { all }, params: {...key},query }})
				return getters['getValidatorDelegations']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryValidatorDelegations API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryValidatorUnbondingDelegations({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.CosmosStakingV1Beta1.query.queryValidatorUnbondingDelegations( key.validator_addr, query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.CosmosStakingV1Beta1.query.queryValidatorUnbondingDelegations( key.validator_addr, {...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'ValidatorUnbondingDelegations', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryValidatorUnbondingDelegations', payload: { options: { all }, params: {...key},query }})
				return getters['getValidatorUnbondingDelegations']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryValidatorUnbondingDelegations API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryDelegation({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.CosmosStakingV1Beta1.query.queryDelegation( key.validator_addr,  key.delegator_addr)).data
				
					
				commit('QUERY', { query: 'Delegation', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryDelegation', payload: { options: { all }, params: {...key},query }})
				return getters['getDelegation']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryDelegation API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryUnbondingDelegation({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.CosmosStakingV1Beta1.query.queryUnbondingDelegation( key.validator_addr,  key.delegator_addr)).data
				
					
				commit('QUERY', { query: 'UnbondingDelegation', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryUnbondingDelegation', payload: { options: { all }, params: {...key},query }})
				return getters['getUnbondingDelegation']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryUnbondingDelegation API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryDelegatorDelegations({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.CosmosStakingV1Beta1.query.queryDelegatorDelegations( key.delegator_addr, query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.CosmosStakingV1Beta1.query.queryDelegatorDelegations( key.delegator_addr, {...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'DelegatorDelegations', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryDelegatorDelegations', payload: { options: { all }, params: {...key},query }})
				return getters['getDelegatorDelegations']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryDelegatorDelegations API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryDelegatorUnbondingDelegations({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.CosmosStakingV1Beta1.query.queryDelegatorUnbondingDelegations( key.delegator_addr, query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.CosmosStakingV1Beta1.query.queryDelegatorUnbondingDelegations( key.delegator_addr, {...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'DelegatorUnbondingDelegations', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryDelegatorUnbondingDelegations', payload: { options: { all }, params: {...key},query }})
				return getters['getDelegatorUnbondingDelegations']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryDelegatorUnbondingDelegations API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryRedelegations({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.CosmosStakingV1Beta1.query.queryRedelegations( key.delegator_addr, query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.CosmosStakingV1Beta1.query.queryRedelegations( key.delegator_addr, {...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'Redelegations', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryRedelegations', payload: { options: { all }, params: {...key},query }})
				return getters['getRedelegations']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryRedelegations API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryDelegatorValidators({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.CosmosStakingV1Beta1.query.queryDelegatorValidators( key.delegator_addr, query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.CosmosStakingV1Beta1.query.queryDelegatorValidators( key.delegator_addr, {...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'DelegatorValidators', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryDelegatorValidators', payload: { options: { all }, params: {...key},query }})
				return getters['getDelegatorValidators']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryDelegatorValidators API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryDelegatorValidator({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.CosmosStakingV1Beta1.query.queryDelegatorValidator( key.delegator_addr,  key.validator_addr)).data
				
					
				commit('QUERY', { query: 'DelegatorValidator', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryDelegatorValidator', payload: { options: { all }, params: {...key},query }})
				return getters['getDelegatorValidator']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryDelegatorValidator API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryHistoricalInfo({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.CosmosStakingV1Beta1.query.queryHistoricalInfo( key.height)).data
				
					
				commit('QUERY', { query: 'HistoricalInfo', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryHistoricalInfo', payload: { options: { all }, params: {...key},query }})
				return getters['getHistoricalInfo']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryHistoricalInfo API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryPool({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.CosmosStakingV1Beta1.query.queryPool()).data
				
					
				commit('QUERY', { query: 'Pool', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryPool', payload: { options: { all }, params: {...key},query }})
				return getters['getPool']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryPool API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryParams({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.CosmosStakingV1Beta1.query.queryParams()).data
				
					
				commit('QUERY', { query: 'Params', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryParams', payload: { options: { all }, params: {...key},query }})
				return getters['getParams']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryParams API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		async sendMsgDelegate({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.CosmosStakingV1Beta1.tx.sendMsgDelegate({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDelegate:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgDelegate:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgUndelegate({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.CosmosStakingV1Beta1.tx.sendMsgUndelegate({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUndelegate:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgUndelegate:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgEditValidator({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.CosmosStakingV1Beta1.tx.sendMsgEditValidator({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgEditValidator:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgEditValidator:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgBeginRedelegate({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.CosmosStakingV1Beta1.tx.sendMsgBeginRedelegate({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgBeginRedelegate:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgBeginRedelegate:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgCreateValidator({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.CosmosStakingV1Beta1.tx.sendMsgCreateValidator({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateValidator:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCreateValidator:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		
		async MsgDelegate({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.CosmosStakingV1Beta1.tx.msgDelegate({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDelegate:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgDelegate:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgUndelegate({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.CosmosStakingV1Beta1.tx.msgUndelegate({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUndelegate:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgUndelegate:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgEditValidator({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.CosmosStakingV1Beta1.tx.msgEditValidator({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgEditValidator:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgEditValidator:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgBeginRedelegate({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.CosmosStakingV1Beta1.tx.msgBeginRedelegate({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgBeginRedelegate:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgBeginRedelegate:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgCreateValidator({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.CosmosStakingV1Beta1.tx.msgCreateValidator({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateValidator:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCreateValidator:Create Could not create message: ' + e.message)
				}
			}
		},
		
	}
}
