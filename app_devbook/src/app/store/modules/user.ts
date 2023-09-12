import { LoginFormDevBookAPI, PostDevBookAPI } from "../../modules/user/interfaces"
import User from "../../modules/user/user"
import { Method, fetchAPIDevBook } from "../../utils/fetch"

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
            commit('SET_USER', user)
        },

        async LogoutAction({ commit }: { commit: any }) {
            //...
        },

        async GetFeed({ commit, state }: { commit: any, state: any }) {
            if (state.user) {
                const feed = await state.user.getUserPosts()
                commit('SET_FEED', feed)
            } else {
                const feed = (await fetchAPIDevBook({ path: '/posts/random', method: Method.GET })).data
                commit('SET_FEED', feed)
            }

        }

    }
})