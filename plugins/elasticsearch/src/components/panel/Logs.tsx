import {
  Alert,
  AlertActionLink,
  AlertVariant,
  Select,
  SelectOption,
  SelectOptionObject,
  SelectVariant,
  Spinner,
} from '@patternfly/react-core';
import { QueryObserverResult, useQuery } from 'react-query';
import React, { useState } from 'react';

import { ILogsData, IQuery } from '../../utils/interfaces';
import { IPluginTimes, PluginCard } from '@kobsio/plugin-core';
import LogsActions from './LogsActions';
import LogsChart from '../panel/LogsChart';
import LogsDocuments from '../panel/LogsDocuments';

interface ILogsProps {
  name: string;
  title: string;
  description?: string;
  queries: IQuery[];
  showChart: boolean;
  times: IPluginTimes;
}

const Logs: React.FunctionComponent<ILogsProps> = ({
  name,
  title,
  description,
  queries,
  showChart,
  times,
}: ILogsProps) => {
  const [showSelect, setShowSelect] = useState<boolean>(false);
  const [selectedQuery, setSelectedQuery] = useState<IQuery>(queries[0]);

  const { isError, isFetching, isLoading, data, error, refetch } = useQuery<ILogsData, Error>(
    ['elasticsearch/logs', selectedQuery, times],
    async () => {
      try {
        const response = await fetch(
          `/api/plugins/elasticsearch/logs/${name}?query=${selectedQuery.query}&timeStart=${times.timeStart}&timeEnd=${times.timeEnd}`,
          {
            method: 'get',
          },
        );
        const json = await response.json();

        if (response.status >= 200 && response.status < 300) {
          return json;
        } else {
          if (json.error) {
            throw new Error(json.error);
          } else {
            throw new Error('An unknown error occured');
          }
        }
      } catch (err) {
        throw err;
      }
    },
    {
      keepPreviousData: true,
    },
  );

  const select = (
    event: React.MouseEvent<Element, MouseEvent> | React.ChangeEvent<Element>,
    value: string | SelectOptionObject,
  ): void => {
    const query = queries.filter((query) => query.name === value);
    if (query.length === 1) {
      setSelectedQuery(query[0]);
    }
    setShowSelect(false);
  };

  return (
    <PluginCard
      title={title}
      description={description}
      actions={<LogsActions name={name} queries={queries} times={times} isFetching={isFetching} />}
    >
      <div>
        {queries.length > 1 ? (
          <div>
            <Select
              variant={SelectVariant.single}
              typeAheadAriaLabel="Select query"
              placeholderText="Select query"
              onToggle={(): void => setShowSelect(!showSelect)}
              onSelect={select}
              selections={selectedQuery.name}
              isOpen={showSelect}
            >
              {queries.map((query, index) => (
                <SelectOption key={index} value={query.name} description={query.query} />
              ))}
            </Select>
            <p>&nbsp;</p>
          </div>
        ) : null}

        {isLoading ? (
          <div className="pf-u-text-align-center">
            <Spinner />
          </div>
        ) : isError ? (
          <Alert
            variant={AlertVariant.danger}
            isInline={true}
            title="Could not get logs"
            actionLinks={
              <React.Fragment>
                <AlertActionLink onClick={(): Promise<QueryObserverResult<ILogsData, Error>> => refetch()}>
                  Retry
                </AlertActionLink>
              </React.Fragment>
            }
          >
            <p>{error?.message}</p>
          </Alert>
        ) : data ? (
          <div>
            {showChart ? (
              <div>
                <LogsChart buckets={data.buckets} />
                <p>&nbsp;</p>
              </div>
            ) : null}

            <LogsDocuments documents={data.documents} fields={selectedQuery.fields} />
          </div>
        ) : null}
      </div>
    </PluginCard>
  );
};

export default Logs;
