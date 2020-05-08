export class UserNotFound extends Error {
    constructor(message: string) {
      super(message); 
      this.name = "UserNotFound";
    }
}
