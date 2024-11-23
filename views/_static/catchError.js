export async function catchError(promise) {
    try {
        const data = await promise;
        return [undefined, data]
    } catch(error) {
        return [error, undefined]
    }
}

