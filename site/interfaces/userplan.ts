export default interface UserPlan {
  id: string;
  createdAt: string;
  updatedAt: string;
  plan_name: string;
  plan_type: string;
  user_id: number;
  plan_id: number;
  active: boolean;
  start_date: string;
  expiry_date: string;
}
