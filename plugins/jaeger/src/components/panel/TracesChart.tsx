import { Card, CardBody } from '@patternfly/react-core';
import { Datum, Node, ResponsiveScatterPlotCanvas, Serie } from '@nivo/scatterplot';
import React, { ReactNode, useMemo } from 'react';
import Trace from './details/Trace';

import { CHART_THEME, ChartTooltip } from '@kobsio/plugin-core';
import { ITrace } from '../../utils/interfaces';
import { doesTraceContainsError } from '../../utils/helpers';

interface IDatum extends Datum {
  label: string;
  size: number;
  trace: ITrace;
}

interface ITracesChartProps {
  name: string;
  traces: ITrace[];
  showDetails?: (details: React.ReactNode) => void;
}

function isIDatum(object: Datum): object is IDatum {
  return (object as IDatum).trace !== undefined;
}

const TracesChart: React.FunctionComponent<ITracesChartProps> = ({ name, traces, showDetails }: ITracesChartProps) => {
  const { series, min, max, first, last } = useMemo<{
    series: Serie[];
    min: number;
    max: number;
    first: number;
    last: number;
  }>(() => {
    // Initialize min and max so that we can simply compare during traversing.
    let minimalSpans = Number.MAX_SAFE_INTEGER;
    let maximalSpans = 0;
    let firstTime = 0;
    let lastTime = 0;
    const result: IDatum[] = [];

    traces.forEach((trace, index) => {
      if (trace.spans.length < minimalSpans) {
        minimalSpans = trace.spans.length;
      }
      if (trace.spans.length > maximalSpans) {
        maximalSpans = trace.spans.length;
      }

      if (trace.startTime < firstTime || index === 0) {
        firstTime = trace.startTime;
      }
      if (trace.startTime > lastTime || index === 0) {
        lastTime = trace.startTime;
      }

      result.push({
        label: trace.traceName,
        size: trace.spans.length,
        trace,
        x: new Date(Math.floor(trace.spans[0].startTime / 1000)),
        y: trace.duration / 1000,
      });
    });

    return {
      first: firstTime,
      last: lastTime,
      max: maximalSpans,
      min: minimalSpans,
      series: [
        {
          data: result.sort((a, b) => (a.x.valueOf() as number) - (b.x.valueOf() as number)),
          id: 'Traces',
        },
      ],
    };
  }, [traces]);

  return (
    <Card isCompact={true}>
      <CardBody>
        <div style={{ height: '250px' }}>
          <ResponsiveScatterPlotCanvas
            axisBottom={{
              format: '%H:%M:%S',
            }}
            axisLeft={null}
            axisRight={null}
            axisTop={null}
            colors={['#0066cc']}
            data={series}
            enableGridX={false}
            enableGridY={false}
            margin={{ bottom: 25, left: 0, right: 0, top: 0 }}
            nodeSize={{ key: 'size', sizes: [15, 75], values: [min, max] }}
            onClick={(node: Node): void => {
              if (showDetails && isIDatum(node.data)) {
                showDetails(<Trace name={name} trace={node.data.trace} close={(): void => showDetails(undefined)} />);
              }
            }}
            // eslint-disable-next-line @typescript-eslint/ban-ts-comment
            // @ts-ignore as the typing expects NodeProps but actually has Node
            renderNode={(ctx: CanvasRenderingContext2D, props: Node): void => {
              // eslint-disable-next-line react/prop-types
              const hasError = isIDatum(props.data) ? doesTraceContainsError(props.data.trace) : false;

              ctx.beginPath();
              // eslint-disable-next-line react/prop-types
              ctx.arc(props.x, props.y, props.size / 2, 0, 2 * Math.PI);
              // eslint-disable-next-line react/prop-types
              ctx.fillStyle = hasError ? '#c9190b' : props.style.color;
              ctx.fill();
            }}
            theme={CHART_THEME}
            tooltip={(tooltip): ReactNode => {
              // eslint-disable-next-line @typescript-eslint/no-explicit-any
              const isFirstHalf = (tooltip.node.data as any).trace.startTime < (last + first) / 2;
              const hasError = isIDatum(tooltip.node.data) ? doesTraceContainsError(tooltip.node.data.trace) : false;
              const squareColor = hasError ? '#c9190b' : '#0066cc';

              return (
                <ChartTooltip
                  anchor={isFirstHalf ? 'right' : 'left'}
                  color={squareColor}
                  label={`${(tooltip.node.data as unknown as IDatum).label} ${tooltip.node.data.formattedY}`}
                  position={[0, 20]}
                  title={tooltip.node.data.formattedX.toString()}
                />
              );
            }}
            xFormat="time:%Y-%m-%d %H:%M:%S.%L"
            xScale={{ type: 'time' }}
            yFormat={(e): string => e + ' ms'}
            yScale={{ max: 'auto', min: 'auto', type: 'linear' }}
          />
        </div>
      </CardBody>
    </Card>
  );
};

export default TracesChart;
