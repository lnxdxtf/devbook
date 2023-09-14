import cdn_paths from '../../cdn_paths.json'
export enum ImageDevBookClass {
    UserProfile = 0,
    Post = 1
}

export async function setImgUrl(imgElementID: string, imgType: ImageDevBookClass | number, imgID: number): Promise<boolean> {
    const imgElement = document.getElementById(imgElementID) as HTMLImageElement | null

    let imgURL = getUrl(imgType, imgID)

    try {
        const response = await fetch(imgURL, { method: 'HEAD', headers: { 'Content-Type': 'image/png' } })
        
        if (response.ok) {
            imgElement!.src = imgURL
            
        } else {
            imgElement!.src = getUrl(imgType)
            if (imgType == ImageDevBookClass.Post) {
                imgElement!.remove()
            }
        }

        return false // loaded
    } catch (err) {
        imgElement!.src = getUrl(imgType)
        return false // not loaded
    }

}

function getUrl(imgType: ImageDevBookClass | number, imgID?: number): string {

    if (imgType === ImageDevBookClass.UserProfile) {
        if (imgID) {
            return `${import.meta.env.VITE_AWS_S3_BUCKET_URI}${cdn_paths.imgs.user.profile}/user_${imgID}.png`
        }
        return `${import.meta.env.VITE_AWS_S3_BUCKET_URI}${cdn_paths.imgs.user.profile}/user_default_0.png`
    }

    if (imgType === ImageDevBookClass.Post) {
        if (imgID) {
            return `${import.meta.env.VITE_AWS_S3_BUCKET_URI}${cdn_paths.imgs.post}/post_${imgID}.png`
        }
        return `${import.meta.env.VITE_AWS_S3_BUCKET_URI}${cdn_paths.imgs.post}/post_default_0.png`
    }

    return `${import.meta.env.VITE_AWS_S3_BUCKET_URI}${cdn_paths.imgs.user.profile}/user_default_0.png`
}
