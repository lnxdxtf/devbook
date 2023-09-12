import { createStore } from "vuex";
import user from './modules/user'
export default createStore({
    state: {
        isMobile: document.body.offsetWidth < 1024
    },
    mutations: {
        SET_ISMOBILE(state, payload) {
            state.isMobile = payload
        }
    },
    modules: {
        user
    }
})