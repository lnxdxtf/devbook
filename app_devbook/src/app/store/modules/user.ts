import { stat } from "fs"
import { LoginFormDevBookAPI, PostDevBookAPI, RandomUserDevBookAPI, RegisterFormDevBookAPI } from "../../modules/user/interfaces"
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
        randomUsers: null,
        loading: false,
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
        },
        SET_RANDOM_USERS(state: any, data: RandomUserDevBookAPI[]) {
            state.randomUsers = data
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

        async LogoutAction({ commit, state }: { commit: any, state: any }): Promise<void> {
            //...
            await state.user.logout()
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

        async GetUserLoggedAction({ commit }: { commit: any }): Promise<void> {
            let user = new User()
            if (user.getLoggedUser()) {
                commit('SET_USER', user)
                commit('SET_AUTHENTICATED', true)
                return
            }
        },

        async GetFeedAction({ commit, state }: { commit: any, state: any }): Promise<void> {
            if (state.authenticated) {
                const feed = await state.user.getUserPosts()
                commit('SET_FEED', feed)
                return
            }
            const feed = (await fetchAPIDevBook({ path: '/posts/random', method: Method.GET }) as ResponseJsonSuccess).data
            commit('SET_FEED', feed)
            return
        },

        async GetRandomUsersAction({ commit }: { commit: any }): Promise<void> {
            const randomUsers = (await fetchAPIDevBook({ path: '/users/random', method: Method.GET }) as ResponseJsonSuccess).data
            commit('SET_RANDOM_USERS', randomUsers)
        },

        async FollowUserAction({ commit, state, dispatch }: { commit: any, state: any, dispatch: any }, userId: number): Promise<void> {
            if (!state.authenticated) return
            await Promise.all([
                state.user.followUser(userId),
                dispatch('GetRandomUsersAction'),
                dispatch('GetFeedAction'),
            ])
        }

    }
})