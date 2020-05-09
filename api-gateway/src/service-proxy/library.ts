import configuration from "../configuration";

const BaseURL = configuration.LIBRARY_URL;

export const LIBRARY_URLS = {
    createBookLoan : `${BaseURL}/api/v1/book-loan/create`,
    bookLoanDetails: (param: string) => `${BaseURL}/api/v1/book-loan/${param}`,
    listBookLoan: `${BaseURL}/api/v1/book-loan`,
    bookLoanAccept:(param: string) => `${BaseURL}/api/v1/book-loan/${param}/accept`,
    bookLoanReject : (param: string) => `${BaseURL}/api/v1/book-loan/${param}/reject`
}