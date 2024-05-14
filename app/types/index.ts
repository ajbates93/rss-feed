export type Feed = {
  ID: number;
  Title: string;
  URL: string;
};

export interface BaseApiResponse {
  success: boolean;
}

export interface AddNewFeedApiResponse extends BaseApiResponse {
  data: Feed;
}
