import { Menu, MenuContent, MenuList } from '@patternfly/react-core';
import React, { memo } from 'react';

import { IPluginPanelProps, IReference, PluginCard, PluginOptionsMissing } from '@kobsio/plugin-core';
import PanelItem from './PanelItem';

interface IPanelProps extends IPluginPanelProps {
  options?: IReference[];
}

// Panel implements the panel component for the dashboards plugin.
export const Panel: React.FunctionComponent<IPanelProps> = ({ defaults, title, description, options }: IPanelProps) => {
  if (!options || !Array.isArray(options)) {
    return (
      <PluginOptionsMissing
        title={title}
        message="Options for Dashboards panel are missing or invalid"
        details="The panel doesn't contain the required options to get dashboards or the provided options are invalid."
        documentation="https://kobs.io/plugins/dashboards"
      />
    );
  }

  return (
    <PluginCard title={title} description={description} transparent={true}>
      <Menu>
        <MenuContent>
          <MenuList>
            {options.map((reference, index) => (
              <PanelItem key={index} defaults={defaults} reference={reference} />
            ))}
          </MenuList>
        </MenuContent>
      </Menu>
    </PluginCard>
  );
};

export default memo(Panel, (prevProps, nextProps) => {
  if (JSON.stringify(prevProps) === JSON.stringify(nextProps)) {
    return true;
  }

  return false;
});
