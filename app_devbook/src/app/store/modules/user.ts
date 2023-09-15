import { LoginFormDevBookAPI, PostDevBookAPI, RegisterFormDevBookAPI } from "../../modules/user/interfaces"
import User from "../../modules/user/user"
import { Method, ResponseJsonSuccess, fetchAPIDevBook } from "../../utils/fetch"

/**
 * This store wraps the User class.
 * The User class wrap all the user data and methods.
 * To Login, Register, Logout, etc, use the actions in this store.
 */
export default ({
    namespaced: true,

    state: () => ({
        user: null,
        feed: null,
        authenticated: false
    }),

    mutations: {
        SET_USER(state: any, data: User) {
            state.user = data
        },
        SET_FEED(state: any, data: PostDevBookAPI[]) {
            state.feed = data
        },
        SET_AUTHENTICATED(state: any, data: boolean) {
            state.authenticated = data
        }
    },

    actions: {
        async LoginAction({ commit }: { commit: any }, loginData: LoginFormDevBookAPI): Promise<boolean> {
            let user = new User()
            if (await user.login(loginData)) {
                commit('SET_USER', user)
                commit('SET_AUTHENTICATED', true)
                return true
            }
            return false
        },

        async LogoutAction({ commit }: { commit: any }) {
            //...
            commit('SET_AUTHENTICATED', false)
            return
        },

        async RegisterAction({ commit }: { commit: any }, registerData: RegisterFormDevBookAPI): Promise<boolean> {
            let user = new User() // Create a new user to use the register method
            if (await user.register(registerData)) {
                return true
            }
            return false
        },

        async GetFeed({ commit, state }: { commit: any, state: any }) {
            if (state.authenticated) {
                const feed = await state.user.getUserPosts()
                commit('SET_FEED', feed)
                return
            }
            const feed = (await fetchAPIDevBook({ path: '/posts/random', method: Method.GET }) as ResponseJsonSuccess).data
            commit('SET_FEED', feed)
            return
        }

    }
})