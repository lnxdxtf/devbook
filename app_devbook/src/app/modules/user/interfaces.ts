export interface PostDevBookAPI {
    id: number,
    title: string,
    content: string,
    author_id: number,
    author_nick: string,
    likes?: number,
    created_at: string | Date,
}

/**
 * This interface is used to define the User class.
 * Login is a required method, but logout and other properties are optional.
 */
export interface UserDevBookAPI {
    id?: number,
    name?: string,
    email?: string,
    token?: string,
    created_at?: string | Date,
}

export interface UserDevBook extends UserDevBookAPI {
    login(form: LoginFormDevBookAPI): Promise<void>,
    logout?(): Promise<void>,
    getUserPosts?(): Promise<PostDevBookAPI[]>,
    editUserPost?(post: PostDevBookAPI): Promise<void>,
    deleteUserPost?(post: PostDevBookAPI): Promise<void>,
}

export interface LoginFormDevBookAPI {
    email: string,
    password: string,
}
