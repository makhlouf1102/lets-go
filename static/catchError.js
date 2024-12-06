export async function catchError(promise) {
    try {
        const data = await promise;
        return [undefined, data];
    } catch (error) {
        console.error("Error occurred:", error);
        return [error, undefined];
    }
}
