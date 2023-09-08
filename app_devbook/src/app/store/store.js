import { createStore } from "vuex";
import user from './modules/user'
export default createStore({
    state: {
        isMobile: document.body.offsetWidth < 1024
    },
    modules: {
        user
    }
})