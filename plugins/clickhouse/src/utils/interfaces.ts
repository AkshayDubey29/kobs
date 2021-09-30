import { IPluginTimes } from '@kobsio/plugin-core';

// IOptions is the interface for all options, which can be set for an ClickHouse query.
export interface IOptions {
  fields?: string[];
  order: string;
  orderBy: string;
  query: string;
  times: IPluginTimes;
}

// IPanelOptions are the options for the panel component of the ClickHouse plugin.
export interface IPanelOptions {
  type: string;
  queries?: IQuery[];
}

export interface IQuery {
  name?: string;
  query?: string;
  fields?: string[];
  order?: string;
  orderBy?: string;
}

// ILogsData is the interface of the data returned from our Go API for the logs view of the ClickHouse plugin.
export interface ILogsData {
  offset: number;
  timeStart: number;
  count?: number;
  took?: number;
  fields?: string[];
  documents?: IDocument[];
  buckets?: IBucket[];
}

export interface IDocument {
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  [key: string]: any;
}

export interface IBucket {
  interval: number;
  count: number;
}

// IDatum, ILabel and IDomain interfaces are used for the logs chart. IDatum is the formate of the data points required
// by '@patternfly/react-charts. ILabel is the formate of the label and IDomain is the formate returned by the
// onBrushDomainChangeEnd function.
export interface IDatum {
  x: Date;
  y: number;
}

export interface ILabel {
  datum: IDatum;
}

export interface IDomain {
  x: Date[];
  y: number[];
}
