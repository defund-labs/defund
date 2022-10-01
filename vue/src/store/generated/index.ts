// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import DefundlabsDefundBroker from './defundlabs.defund.broker'
import DefundlabsDefundEtf from './defundlabs.defund.etf'
import DefundlabsDefundQuery from './defundlabs.defund.query'


export default { 
  DefundlabsDefundBroker: load(DefundlabsDefundBroker, 'defundlabs.defund.broker'),
  DefundlabsDefundEtf: load(DefundlabsDefundEtf, 'defundlabs.defund.etf'),
  DefundlabsDefundQuery: load(DefundlabsDefundQuery, 'defundlabs.defund.query'),
  
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