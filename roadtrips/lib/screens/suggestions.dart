import 'package:flutter/material.dart';
import 'package:roadtrips/screens/search.dart';
import 'package:graphql_flutter/graphql_flutter.dart';

const String suggestionsQuery = """
  query DestinationSuggestions(\$args: RandomDestinationsWithinRing!) {
    randomDestinationsWithinRing(args: \$args) {
      latitude
      longitude
      name
      wazeDeeplink
    }
  }""";

class Destination {
  final double latitude;
  final double longitude;
  final String name;
  final String wazeDeepLink;

  const Destination(
      this.latitude, this.longitude, this.name, this.wazeDeepLink);
}

class SuggestionsScreen extends StatelessWidget {
  static (String, Widget Function(BuildContext)) route =
      ('/destination/suggestions', (context) => const SuggestionsScreen());

  const SuggestionsScreen({super.key});

  @override
  Widget build(BuildContext context) {
    var params = ModalRoute.of(context)!.settings.arguments as SearchParams;

    return Scaffold(
      appBar: AppBar(title: const Text('Suggestions')),
      body: Query(
        options: QueryOptions(document: gql(suggestionsQuery), variables: {
          'args': {
            'filters': const {'type': 'Restaurant'},
            'ring': {
              'center': {
                'latitude': params.ring.center.latitude,
                'longitude': params.ring.center.longitude,
              },
              'innerRadius': {
                'value': params.ring.innerRadius,
                'unit': 'Miles',
              },
              'outerRadius': {
                'value': params.ring.outerRadius,
                'unit': 'Miles',
              },
            }
          }
        }),
        builder: (QueryResult result,
            {VoidCallback? refetch, FetchMore? fetchMore}) {
          if (result.hasException) {
            return Text(result.exception.toString());
          }

          if (result.isLoading) {
            return const Text('Loading');
          }

          List? destinations = result.data?['randomDestinationsWithinRing'];
          if (destinations == null) {
            return const Text('No destinations');
          }

          return ListView(
            children: destinations
                .map((destination) => Destination(
                      destination['latitude'],
                      destination['longitude'],
                      destination['name'],
                      destination['wazeDeeplink'],
                    ))
                .map((destination) => ListTile(
                      title: Text(destination.name),
                      onTap: () {
                        Navigator.pushNamed(context, '/destination/details',
                            arguments: destination);
                      },
                    ))
                .toList(),
          );
        },
      ),
    );
  }
}
