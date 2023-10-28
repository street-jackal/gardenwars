export interface IPlant {
  id: string;
  common?: string;
  botanical?: string;
  height?: (string | null)[];
  characteristics?: string;
  zones?: (string | null)[];
  favorited?: boolean;
}
