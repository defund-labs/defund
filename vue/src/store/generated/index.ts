// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import DefundlabsDefundBroker from './defundlabs.defund.broker'
import DefundlabsDefundEtf from './defundlabs.defund.etf'
import DefundlabsDefundQuery from './defundlabs.defund.query'
import IbcApplicationsInterchainAccountsControllerV1 from './ibc.applications.interchain_accounts.controller.v1'
import IbcApplicationsInterchainAccountsHostV1 from './ibc.applications.interchain_accounts.host.v1'
import IbcApplicationsTransferV1 from './ibc.applications.transfer.v1'
import IbcCoreChannelV1 from './ibc.core.channel.v1'
import IbcCoreClientV1 from './ibc.core.client.v1'
import IbcCoreConnectionV1 from './ibc.core.connection.v1'
import CosmosAuthV1Beta1 from './cosmos.auth.v1beta1'
import CosmosBankV1Beta1 from './cosmos.bank.v1beta1'
import CosmosBaseTendermintV1Beta1 from './cosmos.base.tendermint.v1beta1'
import CosmosCrisisV1Beta1 from './cosmos.crisis.v1beta1'
import CosmosDistributionV1Beta1 from './cosmos.distribution.v1beta1'
import CosmosEvidenceV1Beta1 from './cosmos.evidence.v1beta1'
import CosmosFeegrantV1Beta1 from './cosmos.feegrant.v1beta1'
import CosmosGovV1Beta1 from './cosmos.gov.v1beta1'
import CosmosMintV1Beta1 from './cosmos.mint.v1beta1'
import CosmosParamsV1Beta1 from './cosmos.params.v1beta1'
import CosmosSlashingV1Beta1 from './cosmos.slashing.v1beta1'
import CosmosStakingV1Beta1 from './cosmos.staking.v1beta1'
import CosmosTxV1Beta1 from './cosmos.tx.v1beta1'
import CosmosUpgradeV1Beta1 from './cosmos.upgrade.v1beta1'
import CosmosVestingV1Beta1 from './cosmos.vesting.v1beta1'


export default { 
  DefundlabsDefundBroker: load(DefundlabsDefundBroker, 'defundlabs.defund.broker'),
  DefundlabsDefundEtf: load(DefundlabsDefundEtf, 'defundlabs.defund.etf'),
  DefundlabsDefundQuery: load(DefundlabsDefundQuery, 'defundlabs.defund.query'),
  IbcApplicationsInterchainAccountsControllerV1: load(IbcApplicationsInterchainAccountsControllerV1, 'ibc.applications.interchain_accounts.controller.v1'),
  IbcApplicationsInterchainAccountsHostV1: load(IbcApplicationsInterchainAccountsHostV1, 'ibc.applications.interchain_accounts.host.v1'),
  IbcApplicationsTransferV1: load(IbcApplicationsTransferV1, 'ibc.applications.transfer.v1'),
  IbcCoreChannelV1: load(IbcCoreChannelV1, 'ibc.core.channel.v1'),
  IbcCoreClientV1: load(IbcCoreClientV1, 'ibc.core.client.v1'),
  IbcCoreConnectionV1: load(IbcCoreConnectionV1, 'ibc.core.connection.v1'),
  CosmosAuthV1Beta1: load(CosmosAuthV1Beta1, 'cosmos.auth.v1beta1'),
  CosmosBankV1Beta1: load(CosmosBankV1Beta1, 'cosmos.bank.v1beta1'),
  CosmosBaseTendermintV1Beta1: load(CosmosBaseTendermintV1Beta1, 'cosmos.base.tendermint.v1beta1'),
  CosmosCrisisV1Beta1: load(CosmosCrisisV1Beta1, 'cosmos.crisis.v1beta1'),
  CosmosDistributionV1Beta1: load(CosmosDistributionV1Beta1, 'cosmos.distribution.v1beta1'),
  CosmosEvidenceV1Beta1: load(CosmosEvidenceV1Beta1, 'cosmos.evidence.v1beta1'),
  CosmosFeegrantV1Beta1: load(CosmosFeegrantV1Beta1, 'cosmos.feegrant.v1beta1'),
  CosmosGovV1Beta1: load(CosmosGovV1Beta1, 'cosmos.gov.v1beta1'),
  CosmosMintV1Beta1: load(CosmosMintV1Beta1, 'cosmos.mint.v1beta1'),
  CosmosParamsV1Beta1: load(CosmosParamsV1Beta1, 'cosmos.params.v1beta1'),
  CosmosSlashingV1Beta1: load(CosmosSlashingV1Beta1, 'cosmos.slashing.v1beta1'),
  CosmosStakingV1Beta1: load(CosmosStakingV1Beta1, 'cosmos.staking.v1beta1'),
  CosmosTxV1Beta1: load(CosmosTxV1Beta1, 'cosmos.tx.v1beta1'),
  CosmosUpgradeV1Beta1: load(CosmosUpgradeV1Beta1, 'cosmos.upgrade.v1beta1'),
  CosmosVestingV1Beta1: load(CosmosVestingV1Beta1, 'cosmos.vesting.v1beta1'),
  
}


function load(mod, fullns) {
    return function init(store) {        
        if (store.hasModule([fullns])) {
            throw new Error('Duplicate module name detected: '+ fullns)
        }else{
            store.registerModule([fullns], mod)
            store.subscribe((mutation) => {
                if (mutation.type == 'common/env/INITIALIZE_WS_COMPLETE') {
                    store.dispatch(fullns+ '/init', null, {
                        root: true
                    })
                }
            })
        }
    }
}