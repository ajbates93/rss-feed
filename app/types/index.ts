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

export type MetaApiResponse = {
  page: number;
  totalPages: number;
  limit: number;
  items: number;
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
