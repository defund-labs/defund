import { Client, registry, MissingWalletError } from 'defund-labs-defund-client-ts'

import { IdentifiedClientState } from "defund-labs-defund-client-ts/ibc.core.client.v1/types"
import { ConsensusStateWithHeight } from "defund-labs-defund-client-ts/ibc.core.client.v1/types"
import { ClientConsensusStates } from "defund-labs-defund-client-ts/ibc.core.client.v1/types"
import { ClientUpdateProposal } from "defund-labs-defund-client-ts/ibc.core.client.v1/types"
import { UpgradeProposal } from "defund-labs-defund-client-ts/ibc.core.client.v1/types"
import { Height } from "defund-labs-defund-client-ts/ibc.core.client.v1/types"
import { Params } from "defund-labs-defund-client-ts/ibc.core.client.v1/types"
import { GenesisMetadata } from "defund-labs-defund-client-ts/ibc.core.client.v1/types"
import { IdentifiedGenesisMetadata } from "defund-labs-defund-client-ts/ibc.core.client.v1/types"


export { IdentifiedClientState, ConsensusStateWithHeight, ClientConsensusStates, ClientUpdateProposal, UpgradeProposal, Height, Params, GenesisMetadata, IdentifiedGenesisMetadata };

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
				ClientState: {},
				ClientStates: {},
				ConsensusState: {},
				ConsensusStates: {},
				ConsensusStateHeights: {},
				ClientStatus: {},
				ClientParams: {},
				UpgradedClientState: {},
				UpgradedConsensusState: {},
				
				_Structure: {
						IdentifiedClientState: getStructure(IdentifiedClientState.fromPartial({})),
						ConsensusStateWithHeight: getStructure(ConsensusStateWithHeight.fromPartial({})),
						ClientConsensusStates: getStructure(ClientConsensusStates.fromPartial({})),
						ClientUpdateProposal: getStructure(ClientUpdateProposal.fromPartial({})),
						UpgradeProposal: getStructure(UpgradeProposal.fromPartial({})),
						Height: getStructure(Height.fromPartial({})),
						Params: getStructure(Params.fromPartial({})),
						GenesisMetadata: getStructure(GenesisMetadata.fromPartial({})),
						IdentifiedGenesisMetadata: getStructure(IdentifiedGenesisMetadata.fromPartial({})),
						
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
				getClientState: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ClientState[JSON.stringify(params)] ?? {}
		},
				getClientStates: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ClientStates[JSON.stringify(params)] ?? {}
		},
				getConsensusState: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ConsensusState[JSON.stringify(params)] ?? {}
		},
				getConsensusStates: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ConsensusStates[JSON.stringify(params)] ?? {}
		},
				getConsensusStateHeights: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ConsensusStateHeights[JSON.stringify(params)] ?? {}
		},
				getClientStatus: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ClientStatus[JSON.stringify(params)] ?? {}
		},
				getClientParams: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ClientParams[JSON.stringify(params)] ?? {}
		},
				getUpgradedClientState: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.UpgradedClientState[JSON.stringify(params)] ?? {}
		},
				getUpgradedConsensusState: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.UpgradedConsensusState[JSON.stringify(params)] ?? {}
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
			console.log('Vuex module: ibc.core.client.v1 initialized!')
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
		
		
		
		 		
		
		
		async QueryClientState({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.IbcCoreClientV1.query.queryClientState( key.client_id)).data
				
					
				commit('QUERY', { query: 'ClientState', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryClientState', payload: { options: { all }, params: {...key},query }})
				return getters['getClientState']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryClientState API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryClientStates({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.IbcCoreClientV1.query.queryClientStates(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.IbcCoreClientV1.query.queryClientStates({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'ClientStates', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryClientStates', payload: { options: { all }, params: {...key},query }})
				return getters['getClientStates']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryClientStates API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryConsensusState({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.IbcCoreClientV1.query.queryConsensusState( key.client_id,  key.revision_number,  key.revision_height, query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.IbcCoreClientV1.query.queryConsensusState( key.client_id,  key.revision_number,  key.revision_height, {...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'ConsensusState', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryConsensusState', payload: { options: { all }, params: {...key},query }})
				return getters['getConsensusState']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryConsensusState API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryConsensusStates({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.IbcCoreClientV1.query.queryConsensusStates( key.client_id, query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.IbcCoreClientV1.query.queryConsensusStates( key.client_id, {...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'ConsensusStates', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryConsensusStates', payload: { options: { all }, params: {...key},query }})
				return getters['getConsensusStates']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryConsensusStates API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryConsensusStateHeights({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.IbcCoreClientV1.query.queryConsensusStateHeights( key.client_id, query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.IbcCoreClientV1.query.queryConsensusStateHeights( key.client_id, {...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'ConsensusStateHeights', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryConsensusStateHeights', payload: { options: { all }, params: {...key},query }})
				return getters['getConsensusStateHeights']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryConsensusStateHeights API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryClientStatus({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.IbcCoreClientV1.query.queryClientStatus( key.client_id)).data
				
					
				commit('QUERY', { query: 'ClientStatus', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryClientStatus', payload: { options: { all }, params: {...key},query }})
				return getters['getClientStatus']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryClientStatus API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryClientParams({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.IbcCoreClientV1.query.queryClientParams()).data
				
					
				commit('QUERY', { query: 'ClientParams', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryClientParams', payload: { options: { all }, params: {...key},query }})
				return getters['getClientParams']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryClientParams API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryUpgradedClientState({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.IbcCoreClientV1.query.queryUpgradedClientState()).data
				
					
				commit('QUERY', { query: 'UpgradedClientState', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryUpgradedClientState', payload: { options: { all }, params: {...key},query }})
				return getters['getUpgradedClientState']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryUpgradedClientState API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryUpgradedConsensusState({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.IbcCoreClientV1.query.queryUpgradedConsensusState()).data
				
					
				commit('QUERY', { query: 'UpgradedConsensusState', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryUpgradedConsensusState', payload: { options: { all }, params: {...key},query }})
				return getters['getUpgradedConsensusState']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryUpgradedConsensusState API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
	}
}
