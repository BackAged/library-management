import { GetUserUseCase } from "./getUserUseCase"
import { Gender, User } from "../../entity/user";
import { UserNotFound } from "./errors";

// TODO:-> mocks should be separate folder
let mockUserRepository = {
  createUser: jest.fn(),
  getUserByEmail: jest.fn(),
  listUser: jest.fn(),
  getUser: jest.fn(),
};


describe("Testing GetUserUseCase", () => {
    it("it should return a user", async() => {
        mockUserRepository.createUser.mockClear();
        mockUserRepository.getUserByEmail.mockClear();
        mockUserRepository.listUser.mockClear();

        const user = User.FromUser({
          ID: "sdfsdf",
          age: 2,
          email: "sdfsdf",
          gender: Gender.Male,
          name: "sdfsd",
          password: "234343",
          phone: "2343434"
        });
        mockUserRepository.getUser.mockResolvedValueOnce(user);
        
        const getUserUseCase = new GetUserUseCase(mockUserRepository);
  
        const userReturned = await getUserUseCase.execute(user.ID as string);
  
        expect(userReturned.email).toBe(user.email);
        expect(mockUserRepository.getUser).toHaveBeenCalled();       
    });

    it("it should throw when user not found", async() => {
        mockUserRepository.createUser.mockClear();
        mockUserRepository.getUserByEmail.mockClear();
        mockUserRepository.listUser.mockClear();

        mockUserRepository.getUser.mockResolvedValueOnce(null);
        
        const getUserUseCase = new GetUserUseCase(mockUserRepository);
  
        try {
            const userReturned = await getUserUseCase.execute("blaah");
        } catch(e) {
            expect(e).toBeInstanceOf(UserNotFound)
        }
      
  
        expect(mockUserRepository.getUser).toHaveBeenCalled();       
    });
});