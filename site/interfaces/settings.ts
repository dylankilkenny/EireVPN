export default interface Settings {
  enableCsrf: string;
  enableSubscriptions: string;
  enableAuth: string;
  enableStripe: string;
  authCookieAge: number;
  authCookieName: string;
  authTokenExpiry: number;
  allowedOrigins: string[];
}
