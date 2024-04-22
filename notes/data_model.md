# Data model

## Event
- ID
- Name

### Event/Section
- EventID
- SectionID

### Event/Round
- EventID
- Round
- Date

## Rating System
- ID
- Name

## Player
- ID
- Name

### Player/System
- PlayerID
- SystemID
- Date
- Rating

#### Player/Event/Section/Round
- PlayerID
- EventID
- SectionID
- Round
- BoardNumber
- White
- Black
- Result

