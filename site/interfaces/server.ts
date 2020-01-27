export default interface Server {
  id: string;
  createdAt: string;
  updatedAt: string;
  country: string;
  country_code: number;
  type: string;
  ip: string;
  port: number;
  username: string;
  password: string;
  image_path: string;
}
