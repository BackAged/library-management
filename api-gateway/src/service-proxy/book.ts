import configuration from "../configuration";

const BaseURL = configuration.BOOK_URL;

export const BOOK_URLS = {
    createBook : `${BaseURL}/api/v1/book/create`,
    bookDetails: (param: string) => `${BaseURL}/api/v1/book/${param}`,
    listBook: `${BaseURL}/api/v1/book`,
    listBookByAuthor:(param: string) => `${BaseURL}/api/v1/book/author/${param}`,
    createAuthor : `${BaseURL}/api/v1/author/create`,
    authorDetails: (param :string) => `${BaseURL}/api/v1/author/${param}`,
    listAuthor: `${BaseURL}/api/v1/author`,
}