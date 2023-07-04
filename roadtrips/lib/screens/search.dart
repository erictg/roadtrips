import 'package:flutter/material.dart';

class SearchScreenArgs {
  final String destinationScreenRouteName;

  const SearchScreenArgs(this.destinationScreenRouteName);
}

class Point {
  double latitude;
  double longitude;

  Point(this.latitude, this.longitude);
}

class Ring {
  Point center;
  double innerRadius;
  double outerRadius;

  Ring(this.center, this.innerRadius, this.outerRadius);
}

enum DestinationType { restaurant }

class DestinationTypeFilter {
  DestinationType anyOf;

  DestinationTypeFilter(this.anyOf);
}

class DestinationFilters {
  DestinationTypeFilter? type;

  DestinationFilters(this.type);
}

class SearchParams {
  DestinationFilters? filters;
  Ring ring;

  SearchParams(this.filters, this.ring);
}

class SearchScreen extends StatefulWidget {
  static (String, Widget Function(BuildContext)) route =
      ('/destination/search', (context) => const SearchScreen());

  const SearchScreen({super.key});

  @override
  State<SearchScreen> createState() => _SearchState();
}

class _SearchState extends State<SearchScreen> {
  final _formKey = GlobalKey<FormState>();

  late TextEditingController _ringCenterLatController;
  late TextEditingController _ringCenterLngController;
  late TextEditingController _ringInnerRadiusController;
  late TextEditingController _ringOuterRadiusController;

  @override
  void initState() {
    super.initState();
    _ringCenterLatController = TextEditingController();
    _ringCenterLngController = TextEditingController();
    _ringInnerRadiusController = TextEditingController();
    _ringOuterRadiusController = TextEditingController();
  }

  @override
  Widget build(BuildContext context) {
    final args = ModalRoute.of(context)!.settings.arguments as SearchScreenArgs;

    return Scaffold(
      appBar: AppBar(title: const Text('Search')),
      body: Form(
        key: _formKey,
        child: Column(
          mainAxisAlignment: MainAxisAlignment.spaceAround,
          children: [
            TextFormField(
              decoration: const InputDecoration(
                labelText: 'Ring Center Latitude *',
              ),
              controller: _ringCenterLatController,
              keyboardType: TextInputType.number,
              validator: (value) {
                if (value == null || value.isEmpty) {
                  return 'Ring center latitude is required.';
                }
                return null;
              },
            ),
            TextFormField(
              decoration: const InputDecoration(
                labelText: 'Ring Center Longitude *',
              ),
              controller: _ringCenterLngController,
              keyboardType: TextInputType.number,
              validator: (value) {
                if (value == null || value.isEmpty) {
                  return 'Ring center longitude is required.';
                }
                return null;
              },
            ),
            TextFormField(
              decoration: const InputDecoration(
                labelText: 'Ring Inner Radius *',
                hintText: 'in meters',
              ),
              controller: _ringInnerRadiusController,
              keyboardType: TextInputType.number,
              validator: (value) {
                if (value == null || value.isEmpty) {
                  return 'Inner radius is required.';
                }
                return null;
              },
            ),
            TextFormField(
              decoration: const InputDecoration(
                labelText: 'Ring Outer Radius *',
                hintText: 'in meters',
              ),
              controller: _ringOuterRadiusController,
              keyboardType: TextInputType.number,
              validator: (value) {
                if (value == null || value.isEmpty) {
                  return 'Outer radius is required';
                }
                return null;
              },
            ),
            const DropdownMenu(
              label: Text('Destination Type'),
              enabled: false,
              initialSelection: 'Restaurant',
              dropdownMenuEntries: [
                DropdownMenuEntry(value: 'Restaurant', label: 'Restaurant')
              ],
            ),
            ElevatedButton(
              onPressed: () {
                if (!_formKey.currentState!.validate()) {
                  return;
                }

                var center = Point(double.parse(_ringCenterLatController.text),
                    double.parse(_ringCenterLngController.text));
                var ring = Ring(
                  center,
                  double.parse(_ringInnerRadiusController.text),
                  double.parse(_ringOuterRadiusController.text),
                );
                var filters = DestinationFilters(
                  DestinationTypeFilter(DestinationType.restaurant),
                );

                Navigator.pushNamed(
                  context,
                  args.destinationScreenRouteName,
                  arguments: SearchParams(filters, ring),
                );
              },
              child: const Text('Search'),
            )
          ],
        ),
      ),
    );
  }
}
