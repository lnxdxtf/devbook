/**
 * PostsDevBookAPI is used to define the Post data
 * that comes from the API.
 */
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
    nick?: string,
    email?: string,
    token?: string,
    created_at?: string | Date,
}
/**
 * This interface is used to define a Base User Class in Module User at store.
 */
export interface UserDevBook extends UserDevBookAPI {
    login(form: LoginFormDevBookAPI): Promise<boolean>,
    register?(form: RegisterFormDevBookAPI): Promise<boolean>,
    logout?(): Promise<void>,

    createPost?(post: PostDevBookAPI): Promise<void>,
    editPost?(post: PostDevBookAPI): Promise<void>,
    deletePost?(post: PostDevBookAPI): Promise<void>,
    likePost?(id: number): Promise<void>,
    dislikePost?(id: number): Promise<void>,
    getPosts?(): Promise<PostDevBookAPI[]>,

    follow?(id: number): Promise<void>,
    unfollow?(id: number): Promise<void>,
    followers?(): Promise<any[]>,

}


/**
 * This interface is used to define the form data for login
 */
export interface LoginFormDevBookAPI {
    email: string,
    pswrd: string,
}

/**
 * This interface is used to define the form data for register a new user
 */
export interface RegisterFormDevBookAPI extends LoginFormDevBookAPI {
    pswrd_confirm: string,
    name: string,
    nick: string,
    image: any,
}
