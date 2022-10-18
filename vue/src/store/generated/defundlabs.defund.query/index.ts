import { Client, registry, MissingWalletError } from 'defund-labs-defund-client-ts'

import { Interquery } from "defund-labs-defund-client-ts/defundlabs.defund.query/types"
import { InterqueryResult } from "defund-labs-defund-client-ts/defundlabs.defund.query/types"
import { InterqueryTimeoutResult } from "defund-labs-defund-client-ts/defundlabs.defund.query/types"


export { Interquery, InterqueryResult, InterqueryTimeoutResult };

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
				Interquery: {},
				InterqueryAll: {},
				InterqueryResult: {},
				InterqueryResultAll: {},
				InterqueryTimeoutResult: {},
				InterqueryTimeoutResultAll: {},
				
				_Structure: {
						Interquery: getStructure(Interquery.fromPartial({})),
						InterqueryResult: getStructure(InterqueryResult.fromPartial({})),
						InterqueryTimeoutResult: getStructure(InterqueryTimeoutResult.fromPartial({})),
						
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
				getInterquery: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Interquery[JSON.stringify(params)] ?? {}
		},
				getInterqueryAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.InterqueryAll[JSON.stringify(params)] ?? {}
		},
				getInterqueryResult: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.InterqueryResult[JSON.stringify(params)] ?? {}
		},
				getInterqueryResultAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.InterqueryResultAll[JSON.stringify(params)] ?? {}
		},
				getInterqueryTimeoutResult: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.InterqueryTimeoutResult[JSON.stringify(params)] ?? {}
		},
				getInterqueryTimeoutResultAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.InterqueryTimeoutResultAll[JSON.stringify(params)] ?? {}
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
			console.log('Vuex module: defundlabs.defund.query initialized!')
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
		
		
		
		 		
		
		
		async QueryInterquery({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.DefundlabsDefundQuery.query.queryInterquery( key.storeid)).data
				
					
				commit('QUERY', { query: 'Interquery', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryInterquery', payload: { options: { all }, params: {...key},query }})
				return getters['getInterquery']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryInterquery API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryInterqueryAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.DefundlabsDefundQuery.query.queryInterqueryAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.DefundlabsDefundQuery.query.queryInterqueryAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'InterqueryAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryInterqueryAll', payload: { options: { all }, params: {...key},query }})
				return getters['getInterqueryAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryInterqueryAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryInterqueryResult({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.DefundlabsDefundQuery.query.queryInterqueryResult( key.storeid)).data
				
					
				commit('QUERY', { query: 'InterqueryResult', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryInterqueryResult', payload: { options: { all }, params: {...key},query }})
				return getters['getInterqueryResult']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryInterqueryResult API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryInterqueryResultAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.DefundlabsDefundQuery.query.queryInterqueryResultAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.DefundlabsDefundQuery.query.queryInterqueryResultAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'InterqueryResultAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryInterqueryResultAll', payload: { options: { all }, params: {...key},query }})
				return getters['getInterqueryResultAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryInterqueryResultAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryInterqueryTimeoutResult({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.DefundlabsDefundQuery.query.queryInterqueryTimeoutResult( key.storeid)).data
				
					
				commit('QUERY', { query: 'InterqueryTimeoutResult', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryInterqueryTimeoutResult', payload: { options: { all }, params: {...key},query }})
				return getters['getInterqueryTimeoutResult']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryInterqueryTimeoutResult API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryInterqueryTimeoutResultAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.DefundlabsDefundQuery.query.queryInterqueryTimeoutResultAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.DefundlabsDefundQuery.query.queryInterqueryTimeoutResultAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'InterqueryTimeoutResultAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryInterqueryTimeoutResultAll', payload: { options: { all }, params: {...key},query }})
				return getters['getInterqueryTimeoutResultAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryInterqueryTimeoutResultAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		async sendMsgCreateInterqueryResult({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.DefundlabsDefundQuery.tx.sendMsgCreateInterqueryResult({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateInterqueryResult:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCreateInterqueryResult:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgCreateInterqueryTimeout({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.DefundlabsDefundQuery.tx.sendMsgCreateInterqueryTimeout({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateInterqueryTimeout:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCreateInterqueryTimeout:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgCreateInterquery({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.DefundlabsDefundQuery.tx.sendMsgCreateInterquery({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateInterquery:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCreateInterquery:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		
		async MsgCreateInterqueryResult({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.DefundlabsDefundQuery.tx.msgCreateInterqueryResult({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateInterqueryResult:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCreateInterqueryResult:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgCreateInterqueryTimeout({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.DefundlabsDefundQuery.tx.msgCreateInterqueryTimeout({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateInterqueryTimeout:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCreateInterqueryTimeout:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgCreateInterquery({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.DefundlabsDefundQuery.tx.msgCreateInterquery({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateInterquery:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCreateInterquery:Create Could not create message: ' + e.message)
				}
			}
		},
		
	}
}
