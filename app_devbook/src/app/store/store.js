import { createStore } from "vuex";
import user from './modules/user'
export default createStore({
    state: {
        isMobile: document.body.offsetWidth < 1024,
        devbookCDN: import.meta.env.VITE_AWS_S3_BUCKET_URI
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