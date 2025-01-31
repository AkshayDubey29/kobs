import {
  Button,
  ButtonVariant,
  TextInput,
  Toolbar,
  ToolbarContent,
  ToolbarGroup,
  ToolbarItem,
  ToolbarToggleGroup,
} from '@patternfly/react-core';
import { FilterIcon, SearchIcon } from '@patternfly/react-icons';
import React, { useEffect, useState } from 'react';

import { IOptionsAdditionalFields, Options } from '@kobsio/plugin-core';
import { IOptions } from '../../utils/interfaces';

interface IPageToolbarProps extends IOptions {
  setOptions: (data: IOptions) => void;
}

// PageToolbar is the toolbar for the Elasticsearch plugin page. It allows a user to specify query and to select a start
// time and end time for the query.
const PageToolbar: React.FunctionComponent<IPageToolbarProps> = ({
  query,
  fields,
  times,
  setOptions,
}: IPageToolbarProps) => {
  const [data, setData] = useState<IOptions>({
    query: query,
    times: times,
  });

  // changeQuery changes the value of a query.
  const changeQuery = (value: string): void => {
    setData({ ...data, query: value });
  };

  // onEnter is used to detect if the user pressed the "ENTER" key. If this is the case we are calling the setOptions
  // function to trigger the search.
  // use "SHIFT" + "ENTER".
  const onEnter = (e: React.KeyboardEvent<HTMLInputElement> | undefined): void => {
    if (e?.key === 'Enter' && !e.shiftKey) {
      setOptions({ ...data, fields: fields });
    }
  };

  // changeOptions changes the Elasticsearch option. If the options are changed via the refresh button of the Options
  // component we directly modify the options of the parent component, if not we only change the data of the toolbar
  // component and the user can trigger an action via the search button.
  const changeOptions = (
    refresh: boolean,
    additionalFields: IOptionsAdditionalFields[] | undefined,
    timeEnd: number,
    timeStart: number,
  ): void => {
    const tmpData = { ...data };

    if (refresh) {
      setOptions({
        ...tmpData,
        fields: fields,
        times: { timeEnd: timeEnd, timeStart: timeStart },
      });
    }

    setData({
      ...tmpData,
      times: { timeEnd: timeEnd, timeStart: timeStart },
    });
  };

  useEffect(() => {
    setData({ ...data, query: query, times: times });
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [query, times]);

  return (
    <Toolbar id="elasticsearch-toolbar" style={{ paddingBottom: '0px', zIndex: 300 }}>
      <ToolbarContent style={{ padding: '0px' }}>
        <ToolbarToggleGroup style={{ width: '100%' }} toggleIcon={<FilterIcon />} breakpoint="lg">
          <ToolbarGroup style={{ alignItems: 'flex-start', width: '100%' }}>
            <ToolbarItem style={{ width: '100%' }}>
              <TextInput aria-label="Query" type="text" value={data.query} onChange={changeQuery} onKeyDown={onEnter} />
            </ToolbarItem>
            <ToolbarItem>
              <Options timeEnd={data.times.timeEnd} timeStart={data.times.timeStart} setOptions={changeOptions} />
            </ToolbarItem>
            <ToolbarItem>
              <Button
                variant={ButtonVariant.primary}
                icon={<SearchIcon />}
                onClick={(): void => setOptions({ ...data, fields: fields })}
              >
                Search
              </Button>
            </ToolbarItem>
          </ToolbarGroup>
        </ToolbarToggleGroup>
      </ToolbarContent>
    </Toolbar>
  );
};

export default PageToolbar;
