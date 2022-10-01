import { Client, registry, MissingWalletError } from 'defund-labs-defund-client-ts'

import { Source } from "defund-labs-defund-client-ts/defundlabs.defund.broker/types"
import { Broker } from "defund-labs-defund-client-ts/defundlabs.defund.broker/types"
import { Transfer } from "defund-labs-defund-client-ts/defundlabs.defund.broker/types"
import { BrokerPacketData } from "defund-labs-defund-client-ts/defundlabs.defund.broker/types"
import { NoData } from "defund-labs-defund-client-ts/defundlabs.defund.broker/types"


export { Source, Broker, Transfer, BrokerPacketData, NoData };

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
				Broker: {},
				Brokers: {},
				
				_Structure: {
						Source: getStructure(Source.fromPartial({})),
						Broker: getStructure(Broker.fromPartial({})),
						Transfer: getStructure(Transfer.fromPartial({})),
						BrokerPacketData: getStructure(BrokerPacketData.fromPartial({})),
						NoData: getStructure(NoData.fromPartial({})),
						
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
				getBroker: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Broker[JSON.stringify(params)] ?? {}
		},
				getBrokers: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Brokers[JSON.stringify(params)] ?? {}
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
			console.log('Vuex module: defundlabs.defund.broker initialized!')
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
		
		
		
		 		
		
		
		async QueryBroker({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.DefundlabsDefundBroker.query.queryBroker( key.broker)).data
				
					
				commit('QUERY', { query: 'Broker', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryBroker', payload: { options: { all }, params: {...key},query }})
				return getters['getBroker']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryBroker API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryBrokers({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.DefundlabsDefundBroker.query.queryBrokers(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.DefundlabsDefundBroker.query.queryBrokers({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'Brokers', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryBrokers', payload: { options: { all }, params: {...key},query }})
				return getters['getBrokers']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryBrokers API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		async sendMsgAddConnectionBroker({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.DefundlabsDefundBroker.tx.sendMsgAddConnectionBroker({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgAddConnectionBroker:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgAddConnectionBroker:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgAddLiquiditySource({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.DefundlabsDefundBroker.tx.sendMsgAddLiquiditySource({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgAddLiquiditySource:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgAddLiquiditySource:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		
		async MsgAddConnectionBroker({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.DefundlabsDefundBroker.tx.msgAddConnectionBroker({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgAddConnectionBroker:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgAddConnectionBroker:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgAddLiquiditySource({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.DefundlabsDefundBroker.tx.msgAddLiquiditySource({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgAddLiquiditySource:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgAddLiquiditySource:Create Could not create message: ' + e.message)
				}
			}
		},
		
	}
}
