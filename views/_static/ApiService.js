class ApiService {
    instance = null ;
    
    constructor(){
        if(this.instance != null) {
            this.instance = new ApiService()
        }
    }

    getInstance() {
        return this.instance
    }
}