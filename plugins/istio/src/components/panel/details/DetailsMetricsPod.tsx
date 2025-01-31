import { Alert, AlertActionLink, AlertVariant, Card, CardBody, CardTitle, Spinner } from '@patternfly/react-core';
import { QueryObserverResult, useQuery } from 'react-query';
import React from 'react';

import { Chart, ISeries, convertMetrics } from '@kobsio/plugin-prometheus';
import { IPluginTimes } from '@kobsio/plugin-core';

interface IDetailsMetricsPodProps {
  name: string;
  title: string;
  metric: 'cpu' | 'throttling' | 'memory';
  unit: string;
  namespace: string;
  pod: string;
  times: IPluginTimes;
}

const DetailsMetricsPod: React.FunctionComponent<IDetailsMetricsPodProps> = ({
  name,
  title,
  metric,
  unit,
  namespace,
  pod,
  times,
}: IDetailsMetricsPodProps) => {
  const { isError, isLoading, error, data, refetch } = useQuery<ISeries, Error>(
    ['istio/metricspod', name, metric, namespace, pod, times],
    async () => {
      try {
        const response = await fetch(
          `/api/plugins/istio/metricspod/${name}?timeStart=${times.timeStart}&timeEnd=${times.timeEnd}&metric=${metric}&namespace=${namespace}&pod=${pod}`,
          {
            method: 'get',
          },
        );
        const json = await response.json();

        if (response.status >= 200 && response.status < 300) {
          if (json && json.metrics) {
            return convertMetrics(json.metrics, json.startTime, json.endTime, json.min, json.max);
          } else {
            return { endTime: times.timeEnd, labels: {}, max: 0, min: 0, series: [], startTime: times.timeStart };
          }
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
  );

  return (
    <Card isCompact={true}>
      <CardTitle>{title}</CardTitle>
      <CardBody>
        {isLoading ? (
          <div className="pf-u-text-align-center">
            <Spinner />
          </div>
        ) : isError ? (
          <Alert
            variant={AlertVariant.danger}
            isInline={true}
            title="Could not get metrics"
            actionLinks={
              <React.Fragment>
                <AlertActionLink onClick={(): Promise<QueryObserverResult<ISeries, Error>> => refetch()}>
                  Retry
                </AlertActionLink>
              </React.Fragment>
            }
          >
            <p>{error?.message}</p>
          </Alert>
        ) : data ? (
          <div style={{ height: '300px' }}>
            <Chart
              startTime={data.startTime}
              endTime={data.endTime}
              min={data.min}
              max={data.max}
              options={{ stacked: false, type: 'line', unit: unit }}
              labels={data.labels}
              series={data.series}
            />
          </div>
        ) : null}
      </CardBody>
    </Card>
  );
};

export default DetailsMetricsPod;
