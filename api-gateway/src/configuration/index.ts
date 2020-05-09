interface Config {
    USER_AUTH_URL: string;
    BOOK_URL: string;
    LIBRARY_URL: string;
    APPLICATION_SERVER_PORT: number;
    APP_FORCE_SHUTDOWN_SECOND: number;
}

const config: Config = {
    USER_AUTH_URL: process.env.USER_AUTH_URL || "http://localhost:3000",
    BOOK_URL:  process.env.BOOK_URL || "http://localhost:3001",
    LIBRARY_URL: process.env.LIBRARY_URL || "http://localhost:3002",
    APPLICATION_SERVER_PORT: Number(process.env.APPLICATION_SERVER_PORT) || 8000,
    APP_FORCE_SHUTDOWN_SECOND: Number(process.env.APP_FORCE_SHUTDOWN_SECOND) || 30,
};

export default config;