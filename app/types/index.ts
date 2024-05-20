export type Feed = {
  ID: number;
  Title: string;
  URL: string;
};

export type FeedItem = {
  id: number;
  title: string;
  author: string;
  publishedAt: string;
  url: string;
  imageURL: string;
};

type MetaApiResponse = {
  page: number;
  limit: number;
  total: number;
};

export interface BaseApiResponse {
  success: boolean;
}

export interface AddNewFeedApiResponse extends BaseApiResponse {
  data: Feed;
}

export interface RemoveFeedApiResponse extends BaseApiResponse {}

export interface GetFeedItemsApiResponse extends BaseApiResponse {
  data: FeedItem[];
  meta: MetaApiResponse;
}
