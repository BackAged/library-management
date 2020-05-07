import { LoginUserUseCase } from "./loginUserUseCase"
import { Gender, User } from "../../entity/user";
import { UserNotFound, PasswardMisMatch } from "./errors";

// TODO:-> mocks should be separate folder
let mockUserRepository = {
  createUser: jest.fn(),
  getUserByEmail: jest.fn(),
};
let mockJsonWebTokenManager = {
  generateJSONWebTOken: jest.fn(),
}

describe("Testing LoginUserUseCase", () => {
    it("it should login a user", async() => {
        mockUserRepository.createUser.mockClear();
        mockUserRepository.getUserByEmail.mockClear()

        const user = User.FromUser({
          ID: "sdfsdf",
          age: 2,
          email: "sdfsdf",
          gender: Gender.Male,
          name: "sdfsd",
          password: "234343",
          phone: "2343434"
        })
        mockUserRepository.getUserByEmail.mockResolvedValueOnce(user);
        
        const loginUserUseCase = new LoginUserUseCase(mockUserRepository, mockJsonWebTokenManager);
  
        await loginUserUseCase.execute({
          email: "sdfsdf",
          password: "234343",
        });
  
        expect(mockUserRepository.getUserByEmail).toHaveBeenCalled();
        expect(mockJsonWebTokenManager.generateJSONWebTOken).toHaveBeenCalled();
    });

    it("it should throw error when a user is not found", async() => {
        mockUserRepository.createUser.mockClear();
        mockUserRepository.getUserByEmail.mockClear()

        mockUserRepository.getUserByEmail.mockResolvedValueOnce(null);
        
        const loginUserUseCase = new LoginUserUseCase(mockUserRepository, mockJsonWebTokenManager);
  
        try {
            await loginUserUseCase.execute({
                email: "sdfsdf",
                password: "234343",
            });
        } catch(e) {
            expect(e).toBeInstanceOf(UserNotFound);
        }
        
        expect(mockUserRepository.getUserByEmail).toHaveBeenCalled();
        expect(mockJsonWebTokenManager.generateJSONWebTOken).toHaveBeenCalled();
    });

    it("it should throw when user password doesn't match", async() => {
        mockUserRepository.createUser.mockClear();
        mockUserRepository.getUserByEmail.mockClear()

        const user = User.FromUser({
          ID: "sdfsdf",
          age: 2,
          email: "sdfsdf",
          gender: Gender.Male,
          name: "sdfsd",
          password: "234343",
          phone: "2343434"
        })
        mockUserRepository.getUserByEmail.mockResolvedValueOnce(user);
        
        const loginUserUseCase = new LoginUserUseCase(mockUserRepository, mockJsonWebTokenManager);
  
        try {
            await loginUserUseCase.execute({
                email: "sdfsdf",
                password: "wrong password",
            });
        } catch(e) {
            expect(e).toBeInstanceOf(PasswardMisMatch);
        }
  
        expect(mockUserRepository.getUserByEmail).toHaveBeenCalled();
        expect(mockJsonWebTokenManager.generateJSONWebTOken).toHaveBeenCalled();
    });
});