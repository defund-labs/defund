import { Client, registry, MissingWalletError } from 'defund-labs-defund-client-ts'

import { WeightedVoteOption } from "defund-labs-defund-client-ts/cosmos.gov.v1beta1/types"
import { TextProposal } from "defund-labs-defund-client-ts/cosmos.gov.v1beta1/types"
import { Deposit } from "defund-labs-defund-client-ts/cosmos.gov.v1beta1/types"
import { Proposal } from "defund-labs-defund-client-ts/cosmos.gov.v1beta1/types"
import { TallyResult } from "defund-labs-defund-client-ts/cosmos.gov.v1beta1/types"
import { Vote } from "defund-labs-defund-client-ts/cosmos.gov.v1beta1/types"
import { DepositParams } from "defund-labs-defund-client-ts/cosmos.gov.v1beta1/types"
import { VotingParams } from "defund-labs-defund-client-ts/cosmos.gov.v1beta1/types"
import { TallyParams } from "defund-labs-defund-client-ts/cosmos.gov.v1beta1/types"


export { WeightedVoteOption, TextProposal, Deposit, Proposal, TallyResult, Vote, DepositParams, VotingParams, TallyParams };

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
				Proposal: {},
				Proposals: {},
				Vote: {},
				Votes: {},
				Params: {},
				Deposit: {},
				Deposits: {},
				TallyResult: {},
				
				_Structure: {
						WeightedVoteOption: getStructure(WeightedVoteOption.fromPartial({})),
						TextProposal: getStructure(TextProposal.fromPartial({})),
						Deposit: getStructure(Deposit.fromPartial({})),
						Proposal: getStructure(Proposal.fromPartial({})),
						TallyResult: getStructure(TallyResult.fromPartial({})),
						Vote: getStructure(Vote.fromPartial({})),
						DepositParams: getStructure(DepositParams.fromPartial({})),
						VotingParams: getStructure(VotingParams.fromPartial({})),
						TallyParams: getStructure(TallyParams.fromPartial({})),
						
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
				getProposal: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Proposal[JSON.stringify(params)] ?? {}
		},
				getProposals: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Proposals[JSON.stringify(params)] ?? {}
		},
				getVote: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Vote[JSON.stringify(params)] ?? {}
		},
				getVotes: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Votes[JSON.stringify(params)] ?? {}
		},
				getParams: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Params[JSON.stringify(params)] ?? {}
		},
				getDeposit: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Deposit[JSON.stringify(params)] ?? {}
		},
				getDeposits: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Deposits[JSON.stringify(params)] ?? {}
		},
				getTallyResult: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.TallyResult[JSON.stringify(params)] ?? {}
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
			console.log('Vuex module: cosmos.gov.v1beta1 initialized!')
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
		
		
		
		 		
		
		
		async QueryProposal({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.CosmosGovV1Beta1.query.queryProposal( key.proposal_id)).data
				
					
				commit('QUERY', { query: 'Proposal', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryProposal', payload: { options: { all }, params: {...key},query }})
				return getters['getProposal']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryProposal API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryProposals({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.CosmosGovV1Beta1.query.queryProposals(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.CosmosGovV1Beta1.query.queryProposals({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'Proposals', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryProposals', payload: { options: { all }, params: {...key},query }})
				return getters['getProposals']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryProposals API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryVote({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.CosmosGovV1Beta1.query.queryVote( key.proposal_id,  key.voter)).data
				
					
				commit('QUERY', { query: 'Vote', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryVote', payload: { options: { all }, params: {...key},query }})
				return getters['getVote']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryVote API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryVotes({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.CosmosGovV1Beta1.query.queryVotes( key.proposal_id, query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.CosmosGovV1Beta1.query.queryVotes( key.proposal_id, {...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'Votes', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryVotes', payload: { options: { all }, params: {...key},query }})
				return getters['getVotes']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryVotes API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryParams({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.CosmosGovV1Beta1.query.queryParams( key.params_type)).data
				
					
				commit('QUERY', { query: 'Params', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryParams', payload: { options: { all }, params: {...key},query }})
				return getters['getParams']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryParams API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryDeposit({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.CosmosGovV1Beta1.query.queryDeposit( key.proposal_id,  key.depositor)).data
				
					
				commit('QUERY', { query: 'Deposit', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryDeposit', payload: { options: { all }, params: {...key},query }})
				return getters['getDeposit']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryDeposit API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryDeposits({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.CosmosGovV1Beta1.query.queryDeposits( key.proposal_id, query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.CosmosGovV1Beta1.query.queryDeposits( key.proposal_id, {...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'Deposits', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryDeposits', payload: { options: { all }, params: {...key},query }})
				return getters['getDeposits']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryDeposits API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryTallyResult({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.CosmosGovV1Beta1.query.queryTallyResult( key.proposal_id)).data
				
					
				commit('QUERY', { query: 'TallyResult', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryTallyResult', payload: { options: { all }, params: {...key},query }})
				return getters['getTallyResult']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryTallyResult API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		async sendMsgDeposit({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.CosmosGovV1Beta1.tx.sendMsgDeposit({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDeposit:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgDeposit:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgVote({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.CosmosGovV1Beta1.tx.sendMsgVote({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgVote:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgVote:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgVoteWeighted({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.CosmosGovV1Beta1.tx.sendMsgVoteWeighted({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgVoteWeighted:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgVoteWeighted:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgSubmitProposal({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.CosmosGovV1Beta1.tx.sendMsgSubmitProposal({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSubmitProposal:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgSubmitProposal:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		
		async MsgDeposit({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.CosmosGovV1Beta1.tx.msgDeposit({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDeposit:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgDeposit:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgVote({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.CosmosGovV1Beta1.tx.msgVote({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgVote:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgVote:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgVoteWeighted({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.CosmosGovV1Beta1.tx.msgVoteWeighted({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgVoteWeighted:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgVoteWeighted:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgSubmitProposal({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.CosmosGovV1Beta1.tx.msgSubmitProposal({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSubmitProposal:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgSubmitProposal:Create Could not create message: ' + e.message)
				}
			}
		},
		
	}
}
