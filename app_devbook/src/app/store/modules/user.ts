import { LoginFormDevBookAPI, PostDevBookAPI } from "../../modules/user/interfaces"
import User from "../../modules/user/user"

export default ({
    namespaced: true,

    state: () => ({
        user: null,
        feed: null,
    }),

    mutations: {
        SET_USER(state: any, data: User) {
            state.user = data
        },
        SET_FEED(state: any, data: PostDevBookAPI[]) {
            state.feed = data
        },
    },

    actions: {
        async LoginAction({ commit }: { commit: any, state: any }, loginData: LoginFormDevBookAPI) {
            let user = new User()
            await user.login(loginData)
        },

        async LogoutAction({ commit }: { commit: any }) {
            //...
        },

        async GetFeed({ commit }: { commit: any }) {
            //...   
        }

    }
})